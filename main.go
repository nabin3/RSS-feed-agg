package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

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

	// Defining a router for our server
	mux := http.NewServeMux()

	// "GET /v1/readiness" endpoint
	mux.HandleFunc("GET /v1/readiness", handlerReadiness)

	// "GET /v1/err" endpoint
	mux.HandleFunc("GET /v1/err", handlerError)

	// "POST v1/users" endoint
	mux.HandleFunc("POST /v1/users", cfg.handlerCreateUsers)

	// "GET v1/users" endoint
	mux.HandleFunc("GET /v1/users", cfg.middlewareAuth(handlerGetUserByAPIKey))

	// "POST v1/feeds" endoint
	mux.HandleFunc("POST /v1/feeds", cfg.middlewareAuth(cfg.handlerCreateFeeds))

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
