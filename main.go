package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/nabin3/RSS-feed-agg/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	// Loading environment variable
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldn't load environment")
	}

	// Getting port snumber from .env
	portString := os.Getenv("PORT")

	// Getting databse url from .env
	dbURL := os.Getenv("CONN")

	// Estabilishing connection to databse
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Couldn't estabilish connection with databse")
	}

	// Initializing apiConfig struct
	cfg := &apiConfig{
		DB: database.New(db),
	}

	// Adding our feed scraper
	go cfg.StartScraping(6, 60*time.Second)

	// Defining a router for our server
	mux := http.NewServeMux()

	// "GET /v1/readiness" endpoint
	mux.HandleFunc("GET /v1/readiness", handlerReadiness)

	// "GET /v1/err" endpoint
	mux.HandleFunc("GET /v1/err", handlerError)

	// "POST /v1/users" endoint
	mux.HandleFunc("POST /v1/users", cfg.handlerCreateUsers)

	// "GET /v1/users" endoint
	mux.HandleFunc("GET /v1/users", cfg.middlewareAuth(handlerGetUserByAPIKey))

	// "POST /v1/feeds" endoint
	mux.HandleFunc("POST /v1/feeds", cfg.middlewareAuth(cfg.handlerCreateFeeds))

	// "GET /v1/all-feeds" endpoint
	mux.HandleFunc("GET /v1/all-feeds", cfg.handlerGetAllFeeds)

	// "POST /v1/feed_follows" endpoint
	mux.HandleFunc("POST /v1/feed_follows", cfg.middlewareAuth(cfg.handlerCreateFeedFollow))

	// "DELETE /v1/feed_follows/{feedFollowID}" endpoint
	mux.HandleFunc("DELETE /v1/feed_follows/{feedFollowID}", cfg.handlerDeleteFeedFollow)

	// "GET /v1/feed_follows"
	mux.HandleFunc("GET /v1/feed_follows", cfg.middlewareAuth(cfg.handlerGetAllFeedFollows))

	// "GET /v1/posts/{limit}"
	mux.HandleFunc("GET /v1/posts/{limit}", cfg.middlewareAuth(cfg.HandlerGetPosts))

	// Setting appropriate access control headers
	corsMux := middlewareCors(mux)

	// Defining our server
	ourServer := &http.Server{
		Addr:    "localhost:" + portString,
		Handler: corsMux,
	}

	// Server starts to listen
	fmt.Printf("Listening on Port: %s\n", portString)
	log.Fatal(ourServer.ListenAndServe())
}
