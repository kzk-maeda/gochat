package api

import (
	"encoding/json"
	"fmt"
	"main/data"
	"main/util"
	"net/http"
)

func CreateThread(writer http.ResponseWriter, request *http.Request) {
	session_id, _ := getSessionId(request)
	sess, err := util.APISession(session_id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sess)
		err = request.ParseForm()
		if err != nil {
			fmt.Println("Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			fmt.Println(err, "Cannot get user from session")
		}
		topic := request.Form.Get("topic")
		var conv = data.Thread{}
		if conv, err = user.CreateThread(topic); err != nil {
			fmt.Println("Cannot create thread")
		}
		writer.Header().Set("Content-Type", "application/json")
		output, err := json.MarshalIndent(&conv, "", "\t\t")
		if err != nil {
			fmt.Println(err)
		}
		writer.Write(output)
	}
}

func ListThread(writer http.ResponseWriter, request *http.Request) {
	session_id, _ := getSessionId(request)
	_, err := util.APISession(session_id)
	if err != nil {
		fmt.Println(err)
	} else {
		threads, err := data.Threads()
		if err != nil {
			fmt.Println(err, "Cannot read threads")
		}
		writer.Header().Set("Content-Type", "application/json")
		output, err := json.MarshalIndent(&threads, "", "\t\t")
		if err != nil {
			fmt.Println(err)
		}
		writer.Write(output)
	}
}

func GetThread(writer http.ResponseWriter, request *http.Request) {
	session_id, _ := getSessionId(request)
	sess, err := util.APISession(session_id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sess)
		err = request.ParseForm()
		// URLからthreadのidを取得
		vals := request.URL.Query()
		uuid := vals.Get("id")
		thread, err := data.ThreadByUUID(uuid)
		if err != nil {
			fmt.Println(err, "Cannot read thread")
		} else {
			writer.Header().Set("Content-Type", "application/json")
			output, err := json.MarshalIndent(&thread, "", "\t\t")
			if err != nil {
				fmt.Println(err)
			}
			writer.Write(output)
		}
	}
}

func PostThread(writer http.ResponseWriter, request *http.Request) {

}
