package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/rafiq9090/go-auto-scaling-backend/internal/cache"
	"github.com/rafiq9090/go-auto-scaling-backend/internal/db"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	_ = db.NewPostgres()
	_ = cache.NewRedis()
	mux := http.NewServeMux()

	// Health check (Kubernetes needs this)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Simple api endpoint
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(50 * time.Millisecond)
		w.Write([]byte("Hello from Go Auto Scaling Backend!"))
	})

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	log.Printf("Server running on port %s\n", port)
	log.Fatal(server.ListenAndServe())
}
