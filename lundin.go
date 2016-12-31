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
	server.AddHandlerFrom(goserver.HandlerInfo{"/login", loginGetHandler, loginPostHandler, true})
	server.AddHandlerFrom(goserver.HandlerInfo{"/files/", fileHandler, nil, true})
	server.AddHandler("/sjov/", sjovHandler)
	server.AddHandlerFrom(goserver.HandlerInfo{"/beskeder", nil, threadPostHandler, false})              //TODO
	server.AddHandlerFrom(goserver.HandlerInfo{"/beskeder/", threadGetHandler, nil, false})              //TODO
	server.AddHandlerFrom(goserver.HandlerInfo{"/billeder", imagesGetHandler, imagesPostHandler, false}) //TODO
	server.AddHandlerFrom(goserver.HandlerInfo{"/billeder/", imagesGetHandler, nil, false})              //TODO
	server.AddHandler("/", mainHandler)
	users := loadUsers()
	for _, user := range users {
		server.AddUser(user.Name, user.Password)
	}
	server.ServeOnPort(":80") //93.184.206.223
}

type MainData struct {
	Scripts []string
	User    string
}

func NewMainData(user string, scripts ...string) *MainData {
	return &MainData{scripts, user}
}

func mainHandler(w http.ResponseWriter, r *http.Request, info goserver.Info) {
	data := NewForumData(loadThreads(), info.User())
	temp, err := template.ParseFiles("pages/frontpage.html", "pages/base-start.html", "pages/base-end.html", "pages/header.html", "pages/sidebar.html")
	if err != nil {
		w.Write([]byte("Fejl: " + err.Error()))
	} else {
		temp.Execute(w, data)
	}
}

func sjovHandler(w http.ResponseWriter, r *http.Request, info goserver.Info) {
	data := NewMainData(info.User())
	if len(info.Path) == 0 {
		info.Path = "js"
	}
	temp, err := template.ParseFiles("pages/sjov/"+info.Path+".html", "pages/base-start.html", "pages/base-end.html", "pages/header.html", "pages/kode-sidebar.html")
	if err != nil {
		w.Write([]byte("Fejl: " + err.Error()))
	} else {
		temp.Execute(w, data)
	}
}

func fileHandler(w http.ResponseWriter, r *http.Request, info goserver.Info) {
	if strings.HasSuffix(info.Path, ".css") {
		w.Header().Set("Content-Type", "text/css")
		w.WriteHeader(http.StatusOK)
	} else if info.User() == "" {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}
	file, _ := ioutil.ReadFile("files/" + info.Path)
	w.Write(file)
}

func toJson(thing interface{}) string {
	bytes, err := json.Marshal(thing)
	if err != nil {
		return "{\"error\":\"error\"}"
	}
	return string(bytes)
}

func FromForm(r *http.Request, key string) (string, bool) {
	values, ok := r.Form[key]
	if !ok || len(values) == 0 {
		return "", false
	}
	return values[0], true
}
