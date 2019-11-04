package api

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	User    string
	Threads []string
}

func Index(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "kzk_maeda",
		Threads: []string{"thread1", "thread2", "thread3"},
	}
	json, _ := json.Marshal(post)
	writer.Write(json)
}
