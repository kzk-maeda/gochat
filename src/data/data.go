package data

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "host=127.0.0.1 port=5555 user=root password=password dbname=gochat sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}
