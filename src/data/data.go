package data

import (
	"database/sql"
	"log"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgress", "host=127.0.0.1 port=5555 user=root password=password dbname=gochat sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}
