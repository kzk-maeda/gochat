package main

import (
	"fmt"
	"html/template"
	"main/data"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/public.navbar.html",
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
