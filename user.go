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

func loginHandler(server *goserver.Server, w http.ResponseWriter, r *http.Request, path string, session goserver.Session, user interface{}) {
	if r.Method == "GET" {
		returnLoginPage(w)
		return
	}
	err := r.ParseForm()
	if err != nil {
		returnLoginPage(w)
		return
	}
	users, _ := r.Form["user"]
	passwords, _ := r.Form["password"]
	if len(users) > 0 && len(passwords) > 0 && server.Login(users[0], passwords[0], session) {
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
