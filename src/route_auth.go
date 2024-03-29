package main

import (
	"fmt"
	"net/http"

	"main/data"
	"main/util"
)

// GET /login
func login(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("route /login")
	t := util.ParseTemplateFiles("login", "login.layout", "public.navbar")
	t.Execute(writer, nil)
}

// GET /signup
func signup(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("route /signup")
	util.GenerateHTML(writer, nil, "login.layout", "public.navbar", "signup")
}

// POST signup
func signupAccount(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("route /signup_account")
	err := request.ParseForm()
	if err != nil {
		fmt.Println("err ", err)
	}
	user := data.User{
		Name:     request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}
	fmt.Println("DEBUG : ", request.PostFormValue("name"), request.PostFormValue("email"), request.PostFormValue("password"))
	if err := user.Create(); err != nil {
		fmt.Println("err ", err)
	}
	http.Redirect(writer, request, "/login", 302)
}

// POST /authenticate
func authenticate(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("route /authenticate")
	err := request.ParseForm()
	user, err := data.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		fmt.Println("err ", err)
	}
	if user.Password == data.Encrypt(request.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			fmt.Println("err ", err)
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", 302)
	} else {
		http.Redirect(writer, request, "/login", 302)
	}
}

// GET /logout
func logout(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err != http.ErrNoCookie {
		fmt.Println("Failed to get cookie", err)
		session := data.Session{Uuid: cookie.Value}
		session.DeleteByUUID()
	}
	http.Redirect(writer, request, "/", 302)
}
