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
	server.AddHandler("/js/", jsHandler)
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
