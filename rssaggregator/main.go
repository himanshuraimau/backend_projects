package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, World!")
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT environment variable not set")
	}
	router := chi.NewRouter()
	srv := &http.Server {
		Handler: router,
		Addr:    ":" + portString,
	}
    err := srv.ListenAndServe()

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	fmt.Printf("Listening on port %s\n", portString)
}