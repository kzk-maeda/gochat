package data

import (
	"bytes"
	"database/sql"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	connStr := "host=127.0.0.1 port=5555 user=root password=password dbname=gochat sslmode=disable"
	Db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return
}

// Read SQL file as String
func readSqlFile(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	b := bytes.NewBuffer(content)

	return b.String(), nil
}
