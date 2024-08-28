package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/himanshuraimau/backend_projects/rssaggregator/internal/database"
)

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	   		
            type parameters struct {
				Name string `json:"name"` 
				URL string `json:"url"`

			}
			decoder := json.NewDecoder(r.Body)
			params := parameters{}
			err := decoder.Decode(&params)
			if err != nil {
				respondWithError(w, http.StatusBadRequest, "invalid request payload")
				return
			}
			feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
				ID:        uuid.New(),
				Name:      params.Name,
				URL:       params.URL,
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
                UserID: user.ID,
			})

}