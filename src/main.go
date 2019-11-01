package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("GoChat", version(), "started at", config.Address)

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

	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
