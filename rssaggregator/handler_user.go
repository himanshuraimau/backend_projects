package main

import (
	"net/http"
	"structs"
)

func (apiCfg *apiConfig)handleCreateUser(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w,200,struct{}{})
}