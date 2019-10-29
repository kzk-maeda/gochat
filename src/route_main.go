package main

import (
	"fmt"
	"main/data"
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("route /index")
	threads, err := data.Threads()
	if err != nil {
		fmt.Println("Cannot get threads", err)
	} else {
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, threads, "layout", "public.navbar", "index")
		} else {
			generateHTML(writer, threads, "layout", "private.navbar", "index")
		}
	}
}
