package main

import (
	"fmt"
	"net/http"
)

func httpLog(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header
		len := r.ContentLength
		body := make([]byte, len)
		r.Body.Read(body)
		// fmt.Println(header, string(body))
		fmt.Println("Header: ", header)
		fmt.Println("Body: ", string(body))
		h(w, r)
	}
}
