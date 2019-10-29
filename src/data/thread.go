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

// Create New Thread
func (user *User) CreateThread(topic string) (conv Thread, err error) {
	sql, err := readSqlFile("data/sql/insert_thread.sql")
	if err != nil {
		return
	}
	stmt, err := Db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(createUUID(), topic, user.Id, time.Now()).
		Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	return
}

// Get thread by UUID
func ThreadByUUID(uuid string) (conv Thread, err error) {
	conv = Thread{}
	sql, err := readSqlFile("data/sql/select_thread_by_uuid.sql")
	if err != nil {
		return
	}
	err = Db.QueryRow(sql, uuid).
		Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	return
}

// Get the user who started this thread
func (thread *Thread) User() (user User) {
	user = User{}
	sql, err := readSqlFile("data/sql/select_user_by_id.sql")
	if err != nil {
		return
	}
	Db.QueryRow(sql, thread.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}

// Get thread created data
func (thread *Thread) CreatedAtDate() string {
	conv := Thread{}
	sql, err := readSqlFile("data/sql/select_thread_by_uuid.sql")
	if err != nil {
		return "yyyy/mm/dd"
	}
	err = Db.QueryRow(sql, thread.Uuid).
		Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	return conv.CreatedAt.String()
}

// get the number of posts in a thread
func (thread *Thread) NumReplies() (count int) {
	sql, err := readSqlFile("data/sql/select_count_posts.sql")
	if err != nil {
		return 0
	}
	rows, err := Db.Query(sql, thread.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	rows.Close()
	return
}
