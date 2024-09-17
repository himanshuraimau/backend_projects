package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	_ "github.com/lib/pq" // Import PostgreSQL driver
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/himanshuraimau/backend_projects/rssaggregator/internal/database"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT environment variable not set")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable not set")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Cannot connect to database: ", err)
	}
  db:= database.New(conn)
	apiCfg := apiConfig{
		DB: db ,
	}

	go startScrapping(db, 10,time.Minute)




	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/health", handlerReadiness)
	v1Router.Get("/error", handleErr)
	router.Mount("/v1", v1Router)
	v1Router.Post("/users", apiCfg.handleCreateUser)
	v1Router.Get("/users", apiCfg.authMiddleware(apiCfg.handlerGetUser))
	v1Router.Post("/feed", apiCfg.authMiddleware(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds",apiCfg.handlerGetFeeds)
	v1Router.Post("/feed-follows", apiCfg.authMiddleware(apiCfg.handlerCreateFeedFollow))
	v1Router.Get("/feed-follows", apiCfg.authMiddleware(apiCfg.handlerGetFeedFollows))
	v1Router.Delete("/feed-follows/{feedFollowID}", apiCfg.authMiddleware(apiCfg.handlerDeleteFeedFollow))

	srv := &http.Server{
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