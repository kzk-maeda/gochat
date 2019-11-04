package main

import (
	"fmt"
	"net/http"
	"time"

	"main/api"
	"main/util"
)

func main() {
	fmt.Println("GoChat", util.Version(), "started at", util.Config.Address)

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public/"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", httpLog(index))

	mux.HandleFunc("/login", httpLog(login))
	mux.HandleFunc("/signup", httpLog(signup))
	mux.HandleFunc("/signup_account", httpLog(signupAccount))
	mux.HandleFunc("/authenticate", httpLog(authenticate))
	mux.HandleFunc("/logout", httpLog(logout))

	mux.HandleFunc("/thread/new", httpLog(newThread))
	mux.HandleFunc("/thread/read", httpLog(readThread))
	mux.HandleFunc("/thread/create", httpLog(createThread))
	mux.HandleFunc("/thread/post", httpLog(postThread))

	// API
	mux.HandleFunc("/api/index", httpLog(api.Index))
	mux.HandleFunc("/api/authenticate", httpLog(api.Authenticate))
	mux.HandleFunc("/api/logout", httpLog(api.Logout))

	server := &http.Server{
		Addr:           util.Config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(util.Config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(util.Config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
