package main

import (
	"fmt"
	"net/http"

	"main/data"
)

// GET /login
func login(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("route /login")
	t := parseTemplateFiles("login", "login.layout", "public.navbar")
	t.Execute(writer, nil)
}

// GET /signup
func signup(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("route /signup")
	generateHTML(writer, nil, "login.layout", "public.navbar", "signup")
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
