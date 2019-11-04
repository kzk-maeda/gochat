package main

import (
	"fmt"
	"main/data"
	"main/util"
	"net/http"
)

// GET /thread/new
func newThread(writer http.ResponseWriter, request *http.Request) {
	_, err := util.Session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		util.GenerateHTML(writer, nil, "layout", "private.navbar", "new.thread")
	}
}

// POST /thread/create
func createThread(writer http.ResponseWriter, request *http.Request) {
	sess, err := util.Session(writer, request)
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

// GET /thread/read
func readThread(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	uuid := vals.Get("id")
	thread, err := data.ThreadByUUID(uuid)
	if err != nil {
		fmt.Println(err, " Cannot read thread")
	} else {
		_, err := util.Session(writer, request)
		if err != nil {
			util.GenerateHTML(writer, &thread, "layout", "public.navbar", "public.thread")
		} else {
			util.GenerateHTML(writer, &thread, "layout", "private.navbar", "private.thread")
		}
	}
}

// POST /thread/post
func postThread(writer http.ResponseWriter, request *http.Request) {
	sess, err := util.Session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			fmt.Println("Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			fmt.Println("Cannot get user from session")
		}
		body := request.PostFormValue("body")
		uuid := request.PostFormValue("uuid")
		thread, err := data.ThreadByUUID(uuid)
		if err != nil {
			fmt.Println("Cannot read thread")
		}
		if _, err := user.CreatePost(thread, body); err != nil {
			fmt.Println("Cannot Create Post")
		}
		url := fmt.Sprint("/thread/read?id=", uuid)
		http.Redirect(writer, request, url, 302)
	}
}
