package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/nabin3/RSS-feed-agg/internal/database"
)

// Handler for "POST /v1/feeds"
func (cfg *apiConfig) handlerCreateFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	// Format of recieved data
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
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
		log.Printf("error on handlerCreateFeeds at CreateFeed, details: %v", err)
		respondWithError(w, http.StatusInternalServerError, "couldn't create feed")
		return
	}

	// Creating a new feed_follow for this newly created feed linked to the user_id which was used to create this new feed
	feed_follow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		UserID:    feed.UserID,
		FeedID:    feed.ID,
	})
	if err != nil {
		log.Printf("error on handlerCreateFeeds at CreateFeedFollow, details: %v", err)
		respondWithError(w, http.StatusInternalServerError, "couldn't create feed_follow")
		return
	}
	// resposing with successfully created feed and feed_follow
	respondWithJSON(w, http.StatusOK, databseFeedAndfeedFollowToRespFeedAndFeedFollow(feed, feed_follow))
}

// Handler for "GET /v1/all-feeds"
func (cfg *apiConfig) handlerGetAllFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetAllFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "no feed found")
		return
	}

	// Responding organizely
	respondWithJSON(w, http.StatusOK, allDatabaseFeedToAllFeed(feeds))
}
