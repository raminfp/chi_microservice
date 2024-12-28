package main

import (
	"authentication/cmd/entity"
	"errors"
	"fmt"
	"github.com/go-chi/jwtauth/v5"
	"net/http"
)

var requestPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {

	err := app.Helper.ReadJSON(w, r, &requestPayload)
	if err != nil {
		app.Helper.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	// validate the user against the database
	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		app.Helper.ErrorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		app.Helper.ErrorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}
	// log authentication
	err = app.logRequest("authentication", fmt.Sprintf("%s logged in", user.Email))
	if err != nil {
		app.Helper.ErrorJSON(w, err)
		return
	}
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_email": user.Email})

	payload := entity.JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", requestPayload.Email),
		Data:    tokenString,
	}
	app.Helper.WriteJSON(w, http.StatusAccepted, payload)
}
