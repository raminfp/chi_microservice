package main

import (
	"authentication/cmd/database"
	"authentication/cmd/helper"
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

type Config struct {
	DB     *sql.DB
	Models data.Models
	Helper helper.Ihelper
}

func main() {

	log.Println("Starting authentication service")
	// connect to DB
	conn := database.ConnectToDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}
	// set up config
	app := Config{
		DB:     conn,
		Models: data.New(conn),
		Helper: helper.NewHelper(),
	}
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
