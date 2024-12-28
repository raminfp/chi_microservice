package main

import (
	"errors"
	"net/http"
)

type RequestPayload struct {
	Action string        `json:"action"`
	Auth   AuthPayload   `json:"auth,omitempty"`
	Log    LogPayload    `json:"log,omitempty"`
	Mail   MailPayload   `json:"mail,omitempty"`
	Forget ForgetPayload `json:"forget,omitempty"`
}

type ForgetPayload struct {
	Email string `json:"email"`
}

// HandleSubmission is the main point of entry into the broker. It accepts a JSON
// payload and performs an action based on the value of "action" in that JSON.
func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	switch requestPayload.Action {
	case "auth":
		app.authenticate(w, requestPayload.Auth)
	case "log":
		app.logItem(w, requestPayload.Log)
	case "mail":
		app.sendMail(w, requestPayload.Mail)
	case "forget":
		app.ForgetPass(w, requestPayload.Forget)
	case "profile":
		app.sendMail(w, requestPayload.Mail)
	case "me":
		app.sendMail(w, requestPayload.Mail)
	case "adduser":
		app.sendMail(w, requestPayload.Mail)
	default:
		app.errorJSON(w, errors.New("unknown action"))
	}
}
