package data

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"fmt"
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

func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		fmt.Println("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
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
