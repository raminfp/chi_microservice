package main

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"net/http"
)

//var tokenAuth *jwtauth.JWTAuth

type Response struct {
	StatusCode int
	Msg        string
}

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	// specify who is allowed to connect
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	var res Response
	mux.Use(middleware.Heartbeat("/ping"))
	// OK
	mux.Group(func(r chi.Router) {

		//r.Use(httprate.Limit(
		//	30,             // requests
		//	60*time.Second, // per duration
		//	httprate.WithLimitHandler(app.ResponseMessage),
		//))

		res.Msg = "we have a error"
		res.StatusCode = 404
		data, _ := json.Marshal(res)
		jwtauth.ErrNoTokenFound = errors.New(string(data))
		r.Use(jwtauth.Verifier(app.tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Get("/", app.Home)

	})

	mux.Post("/", app.Home)
	mux.Get("/test", app.AuthTest)
	mux.Get("/logout", app.AuthLogout)
	mux.Post("/handle", app.HandleSubmission)

	return mux
}
