package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rafiq9090/go-auto-scaling-backend/internal/cache"
	"github.com/rafiq9090/go-auto-scaling-backend/internal/db"
	"github.com/rafiq9090/go-auto-scaling-backend/internal/model"
	"github.com/rafiq9090/go-auto-scaling-backend/internal/route"
)

func main() {
	instanceID := os.Getenv("INSTANCE_ID")
	if instanceID == "" {
		instanceID = "unknown"
	}
	log.Printf("Starting server on instance %s\n", instanceID)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	database := db.NewPostgres()
	if err := database.AutoMigrate(&model.Task{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	_ = cache.NewRedis()

	// Initialize Gin router
	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	// Simple api endpoint
	r.GET("/hello", func(c *gin.Context) {
		time.Sleep(50 * time.Millisecond)
		c.String(http.StatusOK, "Hello from Go Auto Scaling Backend! (Instance: %s)", instanceID)
	})

	// Setup Routes
	api := r.Group("/api")
	route.SetupRoute(api)

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	log.Printf("Server running on port %s\n", port)
	log.Fatal(server.ListenAndServe())
}
