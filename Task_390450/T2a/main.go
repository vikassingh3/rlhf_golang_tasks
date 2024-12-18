package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc"

	"example.com/protobuf"
)

const (
	SecretKey             = "your-secret-key"
	Issuer                = "your-issuer"
	Audience              = "your-audience"
	JWTExpirationDuration = time.Hour * 24
)

type authService struct {
	protobuf.UnimplementedAuthServiceServer
}

func (a *authService) ValidateToken(ctx context.Context, req *protobuf.ValidateTokenRequest) (*protobuf.ValidateTokenResponse, error) {
	token, err := jwt.ParseWithClaims(req.Token, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["iss"] != Issuer || claims["aud"] != Audience {
			return nil, fmt.Errorf("invalid issuer or audience")
		}

		user := &protobuf.User{
			Id:       claims["id"].(string),
			Username: claims["username"].(string),
			Email:    claims["email"].(string),
		}
		return &protobuf.ValidateTokenResponse{
			User: user,
		}, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	protobuf.RegisterAuthServiceServer(s, &authService{})

	go func() {
		fmt.Println("gRPC server running on port :50051")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	http.HandleFunc("/validate-token", func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Token not provided", http.StatusUnauthorized)
			return
		}

		response, err := validateToken(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	})

	fmt.Println("HTTP server running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func validateToken(token string) (*protobuf.ValidateTokenResponse, error) {
	service := protobuf.NewAuthServiceClient(&grpc.DialOptions{
		Insecure: true, // Replace with secure credentials in production
	})

	req := &protobuf.ValidateTokenRequest{
		Token: token,
	}

	resp, err := service.ValidateToken(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
