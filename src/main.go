package main

import (
	"fmt"
	"html/template"
	"main/data"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public/"))
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
		"templates/layout.html",
		"templates/navbar.html",
		"templates/index.html",
	}
	templates := template.Must(template.ParseFiles(files...))
	// TODO: threads
	threads, err := data.Threads()
	fmt.Println("threads: ", threads, " err: ", err)
	if err == nil {
		templates.ExecuteTemplate(w, "layout", threads)
	}

	fmt.Println(templates)
}
