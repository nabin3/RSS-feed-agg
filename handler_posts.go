package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/nabin3/RSS-feed-agg/internal/database"
)

// Handler for "GET /v1/posts/{limit}"
func (cfg *apiConfig) HandlerGetPosts(w http.ResponseWriter, r *http.Request, user database.User) {
	limit := r.PathValue("limit")
	// If limit path parameter is not given then we will retrieve only 5 posts
	if limit == "" {
		limit = "5"
	}

	// Parsing integer from given query parameter
	queryLimit, err := strconv.Atoi(limit)
	if err != nil {
		log.Printf("error in handlerGetPosts at strcnv.Atoi, given limit parameter is not suitable to parse integer from it: %v", err)
		respondWithError(w, http.StatusBadRequest, "given limit is not a valid number")
		return
	}

	// Retrieveing posts
	posts, err := cfg.DB.GetPostsByUser(r.Context(), database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  int32(queryLimit),
	})
	if err != nil {
		log.Printf("error in handlerGetPosts at GetPostsByUser: %v", err)
		respondWithError(w, http.StatusInternalServerError, "couldn't retrieve any posts")
	}

	// Successfull response
	respondWithJSON(w, http.StatusOK, allDatabasePostsToPosts(posts))
}
