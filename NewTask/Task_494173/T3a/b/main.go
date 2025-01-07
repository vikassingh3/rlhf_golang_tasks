package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

func connectDB(connString string) (*sql.DB, error) {
	var db *sql.DB
	var err error

	for i := 0; i < 3; i++ { // Retry mechanism
		db, err = sql.Open("postgres", connString)
		if err == nil {
			err = db.Ping()
			if err == nil {
				fmt.Println("Database connection established!")
				return db, nil
			}
		}
		log.Printf("Retrying connection (%d/3): %v", i+1, err)
		time.Sleep(2 * time.Second)
	}
	return nil, fmt.Errorf("failed to connect to database: %v", err)
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Order struct {
	ID     int     `json:"id"`
	UserID int     `json:"user_id"`
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
}

func initJaeger(service string) (opentracing.Tracer, func()) {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "127.0.0.1:6831",
		},
	}

	tracer, closer, err := cfg.New(service, config.Logger(jaeger.StdLogger))
	if err != nil {
		log.Fatalf("Could not initialize Jaeger tracer: %v", err)
	}

	opentracing.SetGlobalTracer(tracer)
	return tracer, func() { closer.Close() }
}

func main() {
	connString := "host=localhost dbname=postgres password=root sslmode=disable"
	db, err := connectDB(connString)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	_, closer := initJaeger("user-service")
	defer closer()

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		span, _ := opentracing.StartSpanFromContext(c.Request.Context(), "request")
		defer span.Finish()
		span.SetTag("service", "user-service")
		c.Request = c.Request.WithContext(opentracing.ContextWithSpan(c.Request.Context(), span))
		c.Next()
	})

	router.POST("/user", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON body"})
			return
		}

		span := opentracing.SpanFromContext(c.Request.Context())
		defer span.Finish()
		span.SetTag("endpoint", "/user")
		span.SetTag("http.method", "POST")

		tx, err := db.Begin()
		if err != nil {
			ext.Error.Set(span, true)
			span.LogKV("error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		defer func() {
			if p := recover(); p != nil {
				tx.Rollback()
				ext.Error.Set(span, true)
				span.LogKV("error", p)
			} else if err != nil {
				tx.Rollback()
				ext.Error.Set(span, true)
				span.LogKV("error", err.Error())
			} else {
				err = tx.Commit()
				if err != nil {
					ext.Error.Set(span, true)
					span.LogKV("error", err.Error())
				}
			}
		}()

		var userID int
		err = tx.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", user.Name, user.Email).Scan(&userID)
		if err != nil {
			ext.Error.Set(span, true)
			span.LogKV("error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user"})
			return
		}

		order := Order{
			UserID: userID,
			Amount: 100.0,
			Status: "pending",
		}

		_, err = tx.Exec("INSERT INTO orders (user_id, amount, status) VALUES ($1, $2, $3)", order.UserID, order.Amount, order.Status)
		if err != nil {
			ext.Error.Set(span, true)
			span.LogKV("error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert order"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"user_id": userID})
	})

	fmt.Println("Starting server on port 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
