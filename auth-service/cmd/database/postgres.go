package database

import (
	"database/sql"
	"log"
	"os"
	"time"
)

var counts int64

func OpenDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func ConnectToDB() *sql.DB {
	// for docker
	dsn := os.Getenv("DSN")
	//dsn := "host=127.0.0.1 port=5432 user=postgres password=password dbname=users sslmode=disable timezone=UTC connect_timeout=5"
	for {
		connection, err := OpenDB(dsn)
		if err != nil {
			log.Println("Postgres not yet ready ...", err)
			counts++
		} else {
			log.Println("Connected to Postgres!")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds....")
		time.Sleep(2 * time.Second)
		continue
	}
}
