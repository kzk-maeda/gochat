package main

import (
	"fmt"
	"html/template"
	"net/http"
	// "github.com/"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templetes/layout.html",
		"templetes/navbar.html",
		"templetes/index.html",
	}
	templates := template.Must(template.ParseFiles(files...))
	// TODO: threads

	fmt.Println(templates)
}
