package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/nabin3/RSS-feed-agg/internal/database"
)

// Handler for "POST /v1/users"
func (cfg *apiConfig) handlerCreateUsers(w http.ResponseWriter, r *http.Request) {
	// Type for recieved data
	type parameters struct {
		Name string `json:"name"`
	}

	// Decoding recieved json data
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("error in handlerCreateUsers func, error dtails: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	// Storing a new user in database
	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		log.Printf("error on handlerCreateUsers, deatils: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}

	// User successfully created
	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}

// Handler for "GET /v1/USERS"
func handlerGetUserByAPIKey(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}
