package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/himanshuraimau/backend_projects/rssaggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {

        type parameters struct {
				Name string `json:"name"`
				Url string `json:"url"`

			}
			decoder := json.NewDecoder(r.Body)
			params := parameters{}
			err := decoder.Decode(&params)
			if err != nil {
				respondWithError(w, http.StatusBadRequest, "invalid request payload")
				return
			}
			feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
				ID:		uuid.New(),
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
				Name: params.Name,
				Url: params.Url,
				UserID: user.ID,

			})
			if err != nil {
				respondWithError(w, http.StatusInternalServerError, "Cannot create feed")
				return
			}
			respondWithJSON(w, 201, databaseFeedToFeed(feed))



}


func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request,) {

	feeds, err := apiCfg.DB.GetFeeds(r.Context() )
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Cannot get feed")
		return
	}
	respondWithJSON(w, 201, databaseFeedsToFeeds(feeds))



}