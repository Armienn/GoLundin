package main

import (
	"html/template"
	"io/ioutil"
	"net/http"

	"encoding/json"

	"strings"

	"github.com/Armienn/GoServer"
)

func main() {
	server := goserver.NewServer(true)
	server.AddHandlerFrom(goserver.HandlerInfo{"/login", loginHandler, true})
	server.AddHandlerFrom(goserver.HandlerInfo{"/files/", fileHandler, true})
	server.AddHandler("/sjov/", sjovHandler)
	server.AddHandler("/", mainHandler)
	users := loadUsers()
	for _, user := range users {
		server.AddUser(user.Name, user.Password)
	}
	server.Serve()
}

type MainData struct {
	Scripts []string
	User    string
}

func NewMainData(user string, scripts ...string) *MainData {
	return &MainData{scripts, user}
}

func mainHandler(server *goserver.Server, w http.ResponseWriter, r *http.Request, path string, session goserver.Session, user interface{}) {
	data := NewMainData(user.(string))
	temp, err := template.ParseFiles("templates/frontpage.html", "templates/base-start.html", "templates/base-end.html", "templates/header.html")
	if err != nil {
		w.Write([]byte("Fejl: " + err.Error()))
	} else {
		temp.Execute(w, data)
	}
}

func sjovHandler(server *goserver.Server, w http.ResponseWriter, r *http.Request, path string, session goserver.Session, user interface{}) {
	data := NewMainData(user.(string), "/files/js/golanguage/golanguage.js")
	temp, err := template.ParseFiles("test.html", "templates/base-start.html", "templates/base-end.html", "templates/header.html")
	if err != nil {
		w.Write([]byte("Fejl: " + err.Error()))
	} else {
		temp.Execute(w, data)
	}
}

func fileHandler(server *goserver.Server, w http.ResponseWriter, r *http.Request, path string, session goserver.Session, user interface{}) {
	if strings.HasSuffix(path, ".css") {
		w.Header().Set("Content-Type", "text/css")
		w.WriteHeader(http.StatusOK)
	} else if user == nil {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}
	file, _ := ioutil.ReadFile("files/" + path)
	w.Write(file)
}

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

func toJson(thing interface{}) string {
	bytes, err := json.Marshal(thing)
	if err != nil {
		return "{\"error\":\"error\"}"
	}
	return string(bytes)
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
