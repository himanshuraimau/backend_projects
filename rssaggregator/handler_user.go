package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/himanshuraimau/backend_projects/rssaggregator/internal/auth"
	"github.com/himanshuraimau/backend_projects/rssaggregator/internal/database"
)

func (apiCfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"` // Corrected struct tag
	}

	var params parameters
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	user, err := apiCfg.DB.CreatedUser(r.Context(), database.CreatedUserParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: sql.NullTime{Time: time.Now().UTC(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now().UTC(), Valid: true},
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Cannot create user")
		return
	}

	respondWithJSON(w, 201, databaseUserToUser(user)) 
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
       		
        apikey,err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, "unauthorized")
			return
		}

		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apikey)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "cannot get user")
			return
		}

		respondWithJSON(w,200,databaseUserToUser(user))


}