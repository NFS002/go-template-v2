package main

import (
	"log"
	"net/http"
	"nfs002/gallery/routing"
	"os"
	"time"

	env "github.com/joho/godotenv"
)

func main() {

	if err := env.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}

	// Initialize router
	router := routing.NewRouter()

	// Create server with timeout
	srv := &http.Server{
		Addr:    ":" + os.Getenv("API_PORT"),
		Handler: router,
		// set timeout due CWE-400 - Potential Slowloris Attack
		ReadHeaderTimeout: 5 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Printf("Failed to start server: %v", err)
	}
}
