package main

import (
	"encoding/json"
	"net/http"
	"github.com/himanshuraimau/backend_projects/rssaggregator/internal/database"
	"github.com/google/uuid"
	"time"
)

func (apiCfg *apiConfig)handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct{
		Name string `name`
	}
	err :=json.NewDecoder(r.Body).Decode(&parameters{})
	if err != nil {
		respondWithError(w,400,"Invalid request payload")
		return
	}
	user,err := api.Cfg.DB.CreateUser(r.Context(),database.CreatedUserParams{
		ID: uuid.New().String(),
		Name: params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),

	})
	if err != nil {
		respondWithError(w,500,"Cannot create user")
		return
	}

	respondWithJSON(w,200,struct{}{})
}