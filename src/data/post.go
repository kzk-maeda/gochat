package data

import (
	"time"
)

// Retrive One Post
func RetrievePost(id int) (post Post, err error) {
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

// Update the post
func (post *Post) UpdatePost(body string) (err error) {
	sql, err := readSqlFile("data/sql/update_post.sql")
	if err != nil {
		return
	}
	_, err = Db.Exec(sql, post.Id, body, time.Now())
	return
}

// Delete the post
func (post *Post) DeletePost() (err error) {
	sql, err := readSqlFile("data/sql/delete_post.sql")
	if err != nil {
		return
	}
	_, err = Db.Exec(sql, post.Id)
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
