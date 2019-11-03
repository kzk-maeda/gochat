package data

import (
	"time"
)

type Post struct {
	Id        int
	Uuid      string
	Body      string
	UserId    int
	ThreadId  int
	CreatedAt time.Time
}

// Retrive One Post
func retrieve(id int) (post Post, err error) {
	post = Post{}
	sql, err := readSqlFile("data/sql/select_post_by_id.sql")
	if err != nil {
		return
	}
	err = Db.QueryRow(sql, id).
		Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt)
	return
}

// Create a new post to a thread
func (user *User) CreatePost(conv Thread, body string) (post Post, err error) {
	sql, err := readSqlFile("data/sql/insert_post.sql")
	if err != nil {
		return
	}
	stmt, err := Db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), body, user.Id, conv.Id, time.Now()).
		Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt)
	return
}

func (post *Post) User() (user User) {
	user = User{}
	sql, err := readSqlFile("data/sql/select_user_by_id.sql")
	if err != nil {
		return
	}
	Db.QueryRow(sql, post.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}
