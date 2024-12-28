package main

import (
	"fmt"
	"github.com/go-chi/jwtauth/v5"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct {
	tokenAuth *jwtauth.JWTAuth
}

func main() {

	app := Config{}
	app.tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
	log.Printf("Starting broker service on port %s\n", webPort)

	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
