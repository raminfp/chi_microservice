package main

import (
	"fmt"
	"github.com/go-chi/jwtauth/v5"
	"net/http"
	"time"
)

func (app *Config) Home(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	fmt.Println(claims)
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}
	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) AuthTest(w http.ResponseWriter, r *http.Request) {
	app.tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, _ := app.tokenAuth.Encode(map[string]interface{}{"user_id": 123, "exp": jwtauth.ExpireIn(2 * time.Minute)})
	payload := jsonResponse{
		Error:   false,
		Message: tokenString,
	}
	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) AuthLogout(w http.ResponseWriter, r *http.Request) {

	_, claims, _ := jwtauth.FromContext(r.Context())
	fmt.Println(claims)

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, _ := app.tokenAuth.Encode(map[string]interface{}{"user_id": 123, "exp": jwtauth.ExpireIn(2 * time.Minute)})
	payload := jsonResponse{
		Error:   false,
		Message: tokenString,
	}
	_ = app.writeJSON(w, http.StatusOK, payload)
}
