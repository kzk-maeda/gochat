package data

import (
	"fmt"
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

// Create a new session for an existing user
func (user *User) CreateSession() (session Session, err error) {
	sql, err := readSqlFile("data/sql/insert_session.sql")
	if err != nil {
		return
	}
	statement := sql
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), user.Email, user.Id, time.Now()).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

//  Get the session for an existing user
func (user *User) Session() (session Session, err error) {
	session = Session{}
	sql, err := readSqlFile("data/sql/select_session.sql")
	if err != nil {
		return
	}
	err = Db.QueryRow(sql, user.Id).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

// Get a single user given the email
func UserByEmail(email string) (user User, err error) {
	user = User{}
	sql, err := readSqlFile("data/sql/select_user_by_email.sql")
	if err != nil {
		return
	}
	err = Db.QueryRow(sql, email).Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

// Get the user from the session
func (session *Session) User() (user User, err error) {
	user = User{}
	sql, err := readSqlFile("data/sql/select_user_by_id.sql")
	err = Db.QueryRow(sql, session.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}

// Create new user
func (user *User) Create() (err error) {
	sql, err := readSqlFile("data/sql/insert_user.sql")
	stmt, err := Db.Prepare(sql)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(createUUID(), user.Name, user.Email, Encrypt(user.Password), time.Now()).
		Scan(&user.Id, &user.Uuid, &user.CreatedAt)
	fmt.Println("create user : ", err)

	return
}
