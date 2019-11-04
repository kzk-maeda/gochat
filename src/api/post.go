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
