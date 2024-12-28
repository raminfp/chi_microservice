package main

import "net/http"

func (app *Config) ResponseMessage(w http.ResponseWriter, r *http.Request) {
	_ = app.writeJSON(w, http.StatusTooManyRequests, "to many request")
}
