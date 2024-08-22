package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
)

func main() {
	fmt.Println("Hello, World!")
	godotenv.Load(".env")
	
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT environment variable not set")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable not set")
	}
    
	conn,err:=sql.Open("postgres", dbURL)
    if err != nil {
		log.Fatal("Cannot connect to database: ", err)
	}


	apiCfg := apiConfig{
		DB: database.New(conn),
	}


	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300, // Maximum value not ignored by any of major browsers
	}))

    v1Router := chi.NewRouter()
	v1Router.Get("/health", handlerReadiness)
	v1Router.Get("/error", handleErr)
	router.Mount("/v1", v1Router)
	v1Router.Post("/users", apiCfg.handleCreateUser)

	srv := &http.Server {
		Handler: router,
		Addr:    ":" + portString,
	}
	log.Println("Starting server at ", portString)
    err = srv.ListenAndServe()

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	fmt.Printf("Listening on port %s\n", portString)
}