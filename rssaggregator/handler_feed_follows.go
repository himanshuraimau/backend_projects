//handler_feed_follows.go

package main

import (
    "encoding/json"
    "net/http"
    "time"
    "fmt"
    "github.com/google/uuid"
    "github.com/himanshuraimau/backend_projects/rssaggregator/internal/database"
		"github.com/go-chi/chi"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
    type parameters struct {
        FeedID uuid.UUID `json:"feed_id"`
    }

    decoder := json.NewDecoder(r.Body)
    params := parameters{}
    err := decoder.Decode(&params)
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Couldn't decode parameters")
        return
    }

    feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
        ID:        uuid.New(),
        CreatedAt: time.Now().UTC(),
        UpdatedAt: time.Now().UTC(),
        FeedID:    params.FeedID,
        UserID:    user.ID,
    })
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Cannot get feed follow: %v", err))
        return
    }

    respondWithJSON(w, http.StatusCreated, databaseFeedFollowToFeedFollow(feedFollow))
}


func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {


	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
			respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Cannot get feed follow: %v", err))
			return
	}

	respondWithJSON(w, http.StatusCreated, databaseFeedFollowsToFeedFollows(feedFollows))
}


func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	 feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	 feedFollowID, err := uuid.Parse(feedFollowIDStr)
	 if err != nil {
		 respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid feed follow ID: %v", err))
		 return
	 }
  err =  apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		 ID: feedFollowID,
		 UserID: user.ID,
	 })
	 if err != nil {
		 respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Cannot delete feed follow: %v", err))
		 return
	 }
	 respondWithJSON(w, http.StatusOK, nil)
}
