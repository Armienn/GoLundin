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
	server.AddGetHandler("/login", loginGetHandler, false)
	server.AddPostHandler("/login", loginPostHandler, false)
	server.AddGetHandler("/static/", staticFileHandler, false)
	server.AddGetHandler("/files/", fileHandler, true)
	server.AddPostHandler("/save/", saveHandler, true)
	server.AddGetHandler("/sjov/", sjovHandler, true)
	server.AddPostHandler("/beskeder", threadPostHandler, true)
	server.AddGetHandler("/beskeder/", threadGetHandler, true)
	server.AddPostHandler("/billeder", imagesPostHandler, true)
	server.AddGetHandler("/billeder", imagesGetHandler, true)
	server.AddGetHandler("/billeder/", imagesGetHandler, true)
	server.AddGetHandler("/", mainHandler, true)
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
	data := struct {
		MainData
		Code string
	}{MainData{nil, info.User()}, "document.getElementById('jsbox').innerHTML = \"Hej du\";"}
	if len(info.Path) == 0 {
		info.Path = "js"
	}
	if strings.HasPrefix(info.Path, "js/") {
		file, _ := ioutil.ReadFile("files/" + info.Path + ".js")
		data.Code = string(file)
		info.Path = "js"
	}
	temp, err := template.ParseFiles("pages/sjov/"+info.Path+".html", "pages/base-start.html", "pages/base-end.html", "pages/header.html", "pages/sidebar-code.html")
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
	}
	file, _ := ioutil.ReadFile("files/" + info.Path)
	w.Write(file)
}

func staticFileHandler(w http.ResponseWriter, r *http.Request, info goserver.Info) {
	if strings.HasSuffix(info.Path, ".css") {
		w.Header().Set("Content-Type", "text/css")
		w.WriteHeader(http.StatusOK)
	}
	file, _ := ioutil.ReadFile("static/" + info.Path)
	w.Write(file)
}

func saveHandler(w http.ResponseWriter, r *http.Request, info goserver.Info) {
	r.ParseForm()
	destination := "/"
	if info.Path == "js" {
		text, ok := FromForm(r, "code")
		if !ok {
			w.Write([]byte("Fejl: Couldn't parse form"))
			return
		}
		name, ok := FromForm(r, "title")
		if !ok {
			w.Write([]byte("Fejl: Couldn't parse form"))
			return
		}
		ioutil.WriteFile("files/js/"+name+".js", []byte(text), 0)
		destination = "/sjov/js/" + name
	}
	http.Redirect(w, r, destination, http.StatusTemporaryRedirect)
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
