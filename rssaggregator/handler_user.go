package main

import (
	"encoding/json"
	"net/http"
	"time"
	"database/sql"

	"github.com/google/uuid"
	"github.com/himanshuraimau/backend_projects/rssaggregator/internal/database"
)

func (apiCfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"` // Corrected struct tag
	}

	var params parameters
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
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

	respondWithJSON(w, http.StatusOK, databaseUserToUser(user)) 
}