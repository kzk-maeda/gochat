package main

import (
	"net/http"
)

func login(writer http.ResponseWriter, request *http.Request) {
	t := parseTemplateFiles("login", "login.layout", "public.navbar")
	t.Execute(writer, nil)
}
