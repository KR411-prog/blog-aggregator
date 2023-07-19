package main

import "net/http"

// now this handler function will have access to Database
func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, struct{}{})
}