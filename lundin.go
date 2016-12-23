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
	server.AddHandler("/beskeder/", threadHandler)
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

type ForumData struct {
	MainData
	Threads []Thread
}

func NewForumData(threads []Thread, user string, scripts ...string) *ForumData {
	return &ForumData{MainData{scripts, user}, threads}
}

func NewMainData(user string, scripts ...string) *MainData {
	return &MainData{scripts, user}
}

func mainHandler(server *goserver.Server, w http.ResponseWriter, r *http.Request, path string, session goserver.Session, user interface{}) {
	data := NewForumData(loadThreads(), user.(string))
	temp, err := template.ParseFiles("pages/frontpage.html", "pages/base-start.html", "pages/base-end.html", "pages/header.html")
	if err != nil {
		w.Write([]byte("Fejl: " + err.Error()))
	} else {
		temp.Execute(w, data)
	}
}

func sjovHandler(server *goserver.Server, w http.ResponseWriter, r *http.Request, path string, session goserver.Session, user interface{}) {
	data := NewMainData(user.(string), "/files/js/golanguage/golanguage.js")
	if len(path) == 0 {
		path = "golanguage"
	}
	temp, err := template.ParseFiles("pages/"+path+".html", "pages/base-start.html", "pages/base-end.html", "pages/header.html", "pages/kode-header.html")
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

func toJson(thing interface{}) string {
	bytes, err := json.Marshal(thing)
	if err != nil {
		return "{\"error\":\"error\"}"
	}
	return string(bytes)
}