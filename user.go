package main

import (
	"html/template"
	"io/ioutil"
	"net/http"

	"encoding/json"

	"github.com/Armienn/GoServer"
)

type User struct {
	Name     string
	Password string
}

func loadUsers() []User {
	file, err := ioutil.ReadFile("users.json")
	if err != nil {
		panic(err)
	}
	var users []User
	err = json.Unmarshal(file, &users)
	if err != nil {
		panic(err)
	}
	return users
}

func loginGetHandler(w http.ResponseWriter, r *http.Request, info goserver.Info) {
	returnLoginPage(w)
}

func loginPostHandler(w http.ResponseWriter, r *http.Request, info goserver.Info) {
	err := r.ParseForm()
	if err != nil {
		returnLoginPage(w)
		return
	}
	users, _ := r.Form["user"]
	passwords, _ := r.Form["password"]
	if len(users) > 0 && len(passwords) > 0 && info.Server.Login(users[0], passwords[0], info.Session) {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	returnLoginPage(w)
}

func returnLoginPage(w http.ResponseWriter) {
	temp, err := template.ParseFiles("login.html")
	if err != nil {
		w.Write([]byte("Fejl: " + err.Error()))
	} else {
		temp.Execute(w, nil)
	}
}
