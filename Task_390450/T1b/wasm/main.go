
package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/golang-jwt/jwt"
    "github.com/golang/protobuf/ptypes/empty"
    "github.com/golang/protobuf/ptypes/timestamp"
    "google.golang.org/grpc"

    "example.com/protobuf"
)

const (
    // SecretKey is the secret key for JWT signing.
    SecretKey = "your-secret-key"
    // JWTExpirationDuration is the duration for which JWT tokens are valid.
    JWTExpirationDuration = time.Hour * 24
)

type authService struct {
    protobuf.UnimplementedAuthServiceServer
}

func (a *authService) Login(ctx context.Context, req *protobuf.LoginRequest) (*protobuf.LoginResponse, error) {
    // Implement your login logic here
    // For demonstration purposes, we'll just return a dummy user

    user := &protobuf.User{
        Id:        "1",
        Username:  "user",
        Email:     "user@example.com",
        CreatedAt: &timestamp.Timestamp{Seconds: time.Now().Unix()},
    }

    token, err := generateJWT(user)
    if err != nil {
        return nil, err
    }

    return &protobuf.LoginResponse{
        Token: token,
        User:   user,
    }, nil
}

func (a *authService) ValidateToken(ctx context.Context, req *protobuf.ValidateTokenRequest) (*protobuf.ValidateTokenResponse, error) {
    token, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(SecretKey), nil
    })

    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
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

func generateJWT(user *protobuf.User) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id":       user.Id,
        "username": user.Username,
        "email":    user.Email,
        "exp":      time.Now().Add(JWTExpirationDuration).Unix(),
    })

    return token.SignedString([]byte(SecretKey))
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    protobuf.RegisterAuthServiceServer(s, &authService{})

    // Start the gRPC server
    go func() {
        fmt.Println("gRPC server running on port :50051")
        if err := s.Serve(lis); err != nil {
            log.Fatalf("failed to serve: %v", err)
        }
    }()

    // Start the HTTP server for WebAssembly serving
    fs := http.FileServer(http.Dir("webassembly"))
    http.Handle("/", fs)
    fmt.Println("WebAssembly server running on port :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
