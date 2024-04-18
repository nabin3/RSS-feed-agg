package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/nabin3/RSS-feed-agg/internal/database"
)

func (cfg *apiConfig) handlerCreateFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	// Format of recieved data
	type parameters struct {
		Name string
		Url  string
	}

	// Decoding recieved data
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("error in handlerCreateFeeds, details: %v", err)
		respondWithError(w, http.StatusInternalServerError, "couldn't decode parameters")
		return
	}

	// Creating new feed
	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		log.Printf("error on handlerCreateFeeds, details: %v", err)
		respondWithError(w, http.StatusInternalServerError, "couldn't create feed")
		return
	}

	// resposing with successfully created feed
	respondWithJSON(w, http.StatusOK, databaseFeedToFeed(feed))
}
