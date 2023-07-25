package main

import (
	"blogaggregator/internal/auth"
	"blogaggregator/internal/database"
	"fmt"
	"net/http"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User) //this includes 3rd param. but
//it does not match the http handler

// so we write middlewareAuth to take authedHandler and return HandlerFunc
func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400,fmt.Sprintf("Couldnt get user: %v", err))
			return
		}

	handler(w,r,user)
	}
}
