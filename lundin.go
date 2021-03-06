package main

import (
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
	server.AddGetHandler("/files/", fileHandler, true)
	server.AddPostHandler("/save/", saveHandler, true)
	server.AddPostHandler("/billeder", imagesPostHandler, true)
	ThreadAPI(server, "/beskeder/")
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
	file, _ := ioutil.ReadFile("index.html")
	w.Write(file)
}

func fileHandler(w http.ResponseWriter, r *http.Request, info goserver.Info) {
	if strings.HasSuffix(info.Path, ".css") {
		w.Header().Set("Content-Type", "text/css")
		w.WriteHeader(http.StatusOK)
	}
	file, _ := ioutil.ReadFile("files/" + info.Path)
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
