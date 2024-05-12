package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/nabin3/RSS-feed-agg/internal/database"
)

// Handler for "POST /v1/feed_follows"
func (cfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	// Format of request body
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	// Decoding recieved data
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("error in handlerCreateFeedFollow at decoding request data, details: %v", err)
		respondWithError(w, http.StatusInternalServerError, "couldn't create feed-follow for requested feed-id")
		return
	}

	// inserting a neew feed_follow in database
	feed_follow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		log.Printf("error in handlerCreateFeedFollow at CreateFeedFollow, details %v", err)
		respondWithError(w, http.StatusInternalServerError, "couldn't create feed-follow for requested feed-id")
		return
	}

	// Successfull response
	respondWithJSON(w, http.StatusOK, databaseFeedFollowToRespFeedFollow(feed_follow))
}

// Handler for "DELETE /v1/feed_follows/{feedFollowID}"
func (cfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request) {
	// Extractring query parameter
	ID := r.PathValue("feedFollowID")

	// Parsing FeedFollowID as UUID from feedFollowID query p[arameter which is a string
	feedFollowID, err := uuid.Parse(ID)
	if err != nil {
		log.Printf("error in handlerDeleteFeedFollow at parsing feedFollowId, details %v", err)
		respondWithError(w, http.StatusBadRequest, "invalid feedFollowID")
		return
	}

	// Deleting a feed-follow
	err = cfg.DB.DeleteFeedFollow(r.Context(), feedFollowID)
	if err != nil {
		log.Printf("error in handlerDeleteFeedFollow at DeleteFeedFollow, detail %v", err)
		respondWithError(w, http.StatusBadRequest, "couldn't delete mentioned feed-follow")
	}

	// Successfull response
	respondWithJSON(w, http.StatusOK, "successfully deleted mentioned feed-follow")
}

// Handler for "GET /v1/feed_follows"
func (cfg *apiConfig) handlerGetAllFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	// Retrieveing all feed_follows for a given user_id
	allFeedFollows, err := cfg.DB.GetFeedFollowsForUser(r.Context(), user.ID)
	if err != nil {
		log.Printf("error in handlerGetAllFeedFollows at GetFeedFollowsForUser, detail %v", err)
		respondWithError(w, http.StatusBadRequest, "couldn't get feed-follows")
		return
	}

	// Successfull response
	respondWithJSON(w, http.StatusOK, allDatabaseFeedFollowsToAllFeedFollowsOfUser(allFeedFollows))
}
