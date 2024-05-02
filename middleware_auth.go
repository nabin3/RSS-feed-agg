package main

import (
	"net/http"

	"github.com/nabin3/RSS-feed-agg/internal/auth"
	"github.com/nabin3/RSS-feed-agg/internal/database"
)

// Giving signature of handler which can check for authorization
type authHandler func(http.ResponseWriter, *http.Request, database.User)

// Middleware function for handler which will check for authentication
func (cfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "couldn't find api_key")
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "couldn't find user")
			return
		}

		handler(w, r, user)
	}
}
