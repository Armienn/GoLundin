package main

import (
	"html/template"
	"io/ioutil"
	"net/http"

	"encoding/json"

	"github.com/Armienn/GoServer"
)

func main() {
	server := goserver.NewServer(true)
	server.AddHandlerFrom(goserver.HandlerInfo{"/login", loginHandler, true})
	server.AddHandler("/js/", jsHandler)
	server.AddHandlerFrom(goserver.HandlerInfo{"/css/", cssHandler, true})
	server.AddHandler("/", viewHandler)
	users := loadUsers()
	for _, user := range users {
		server.AddUser(user.Name, user.Password)
	}
	server.Serve()
}

func viewHandler(server *goserver.Server, w http.ResponseWriter, r *http.Request, path string, session goserver.Session, user interface{}) {
	data := struct{ Count int }{0}
	value, ok := session.Get("musle")
	if ok {
		data.Count = value.(int)
	}
	data.Count++
	session.Set("musle", data.Count)
	temp, err := template.ParseFiles("test.html")
	if err != nil {
		w.Write([]byte("Fejl: " + err.Error()))
	} else {
		temp.Execute(w, data)
	}
}

func jsHandler(server *goserver.Server, w http.ResponseWriter, r *http.Request, path string, session goserver.Session, user interface{}) {
	file, _ := ioutil.ReadFile("gopher/" + path)
	w.Write(file)
}

func cssHandler(server *goserver.Server, w http.ResponseWriter, r *http.Request, path string, session goserver.Session, user interface{}) {
	w.Header().Set("Content-Type", "text/css")
	w.WriteHeader(http.StatusOK)
	file, _ := ioutil.ReadFile("css/" + path)
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
