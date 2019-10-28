package main

import (
	"fmt"
	"net/http"
)

// GET /thread/new
func newThread(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		generateHTML(writer, nil, "layout", "private.navbar", "new.thread")
	}
}

// POST /thread/create
func createThread(writer http.ResponseWriter, request *http.Request) {
	sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		fmt.Println(sess)
		err = request.ParseForm()
		if err != nil {
			// danger(err, "Cannot parse form")
			fmt.Println("Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			// danger(err, "Cannot get user from session")
			fmt.Println(err, "Cannot get user from session")
		}
		topic := request.PostFormValue("topic")
		if _, err := user.CreateThread(topic); err != nil {
			// danger(err, "Cannot create thread")
			fmt.Println("Cannot create thread")
		}
		http.Redirect(writer, request, "/", 302)
	}
}
