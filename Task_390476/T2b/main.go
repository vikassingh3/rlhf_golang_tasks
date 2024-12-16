package main

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid" // Correct UUID generation package
	// "github.com/rs/zerolog"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
)

// RequestInfo holds information about a request
type RequestInfo struct {
	RequestID string `json:"request_id"`
	Method    string `json:"method"`
	URL       string `json:"url"`
}

// ResponseInfo holds information about a response
type ResponseInfo struct {
	StatusCode int    `json:"status_code"`
	Body       string `json:"body"`
}

func init() {
	// Configure zerolog
	zerolog.TimeFieldFormat = "2006-01-02T15:04:05.000Z"
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	hlog.RegisterDefault(&logger)
}

func generateRequestID() string {
	return uuid.New().String() // Correct UUID generation
}

// Custom ResponseWriter to capture response status and body
type responseRecorder struct {
	http.ResponseWriter
	statusCode int
	body       bytes.Buffer
}

func (r *responseRecorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func (r *responseRecorder) Write(p []byte) (n int, err error) {
	n, err = r.body.Write(p)
	if err == nil {
		err = r.ResponseWriter.Write(p)
	}
	return n, err
}

func tokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := &responseRecorder{ResponseWriter: w}
		ctx := r.Context()
		requestID := generateRequestID()
		ctx = context.WithValue(ctx, "request_id", requestID)

		// Log request start
		hlog.FromContext(ctx).Info().
			Str("request_id", requestID).
			Msgf("Request started: %s %s", r.Method, r.URL.String())

		// Simulate token handling
		token := r.Header.Get("Authorization")
		if token == "" {
			hlog.FromContext(ctx).Warn().
				Str("request_id", requestID).
				Msg("Token not provided")
			http.Error(w, "Token required", http.StatusUnauthorized)
			return
		}

		hlog.FromContext(ctx).Info().
			Str("request_id", requestID).
			Str("token", token).
			Msg("Token validated")

		// Proceed to the next handler
		next.ServeHTTP(rec, r.WithContext(ctx))

		// Log request end and response info
		response := ResponseInfo{
			StatusCode: rec.statusCode,
			Body:       rec.body.String(),
		}

		hlog.FromContext(ctx).Info().
			Str("request_id", requestID).
			Interface("response", response).
			Msg("Request completed")
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := ctx.Value("request_id").(string)

	hlog.FromContext(ctx).Info().
		Str("request_id", requestID).
		Msg("Handling request")

	w.Write([]byte(`{"message": "Hello, world!"}`))
}

func main() {
	// Register the handler and middleware
	http.HandleFunc("/", tokenMiddleware(http.HandlerFunc(handler)))

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
