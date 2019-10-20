package data

import (
	"time"
)

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

// Get All Threads in the database and returns it
func Threads() (threads []Thread, err error) {
	sql, err := readSqlFile("data/sql/select_threads.sql")
	if err != nil {
		return nil, err
	}
	rows, err := Db.Query(sql)
	if err != nil {
		return
	}
	for rows.Next() {
		conv := Thread{}
		if err = rows.Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt); err != nil {
			return
		}
		threads = append(threads, conv)
	}
	rows.Close()
	return
}
