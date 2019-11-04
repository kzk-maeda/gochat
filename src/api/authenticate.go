package api

import (
	"encoding/json"
	"fmt"
	"main/data"
	"net/http"
	// "main/util"
)

// GET authenticate
func Authenticate(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	email := request.Form.Get("email")
	password := request.Form.Get("password")
	fmt.Println(email, password)
	// emailでUserを検索
	user, err := data.UserByEmail(email)
	if err != nil {
		fmt.Println(email, " is not found in database. ", err)
	}
	var session data.Session
	if user.Password == data.Encrypt(password) {
		session, err = user.CreateSession()
		if err != nil {
			fmt.Println(err)
		}
		uuid := session.Uuid
		fmt.Println("UUID : ", uuid)
	} else {
		err := "password invalid"
		fmt.Println(err)
	}
	// sessionをreturn
	writer.Header().Set("Content-Type", "application/json")
	output, err := json.MarshalIndent(&session, "", "\t\t")
	if err != nil {
		fmt.Println(err)
	}
	writer.Write(output)
}
