package main

import (
	"fmt"
	"main/data"
	"main/util"
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("route /index")
	threads, err := data.Threads()
	if err != nil {
		fmt.Println("Cannot get threads", err)
	} else {
		_, err := util.Session(writer, request)
		if err != nil {
			util.GenerateHTML(writer, threads, "layout", "public.navbar", "index")
		} else {
			util.GenerateHTML(writer, threads, "layout", "private.navbar", "index")
		}
	}
}
