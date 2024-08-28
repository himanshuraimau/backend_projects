package main

import (
	"net/http"

	"github.com/himanshuraimau/backend_projects/rssaggregator/internal/auth"
	"github.com/himanshuraimau/backend_projects/rssaggregator/internal/database"
)

// Define the authHandler type with appropriate parameters
type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) authMiddleware(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusForbidden, "unauthorized")
			return
		}

		user, err := cfg.DB.GetUserByApiKey(r.Context(), apikey)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "cannot get user")
			return
		}

		// Pass the user to the handler
		handler(w, r, user)
	}
}
