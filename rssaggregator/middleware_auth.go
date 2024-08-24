package main

import (
	"net/http"

	"github.com/himanshuraimau/backend_projects/rssaggregator/internal/database"
)

func authHandler (http.ResponseController,*http.Request,database.Queries)


func (cfg *apiConfig) authMiddleware(handler authHandler) http.HandlerFunc {
	   return func(w http.ResponseWriter, r *http.Request) {
	    apikey,err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, "unauthorized")
			return
		}

		user,err := apiCfg.DB.GetUserByApiKey(r.Context(),apikey)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "cannot get user")
			return
		}
		handler(w,r,user)

	
}