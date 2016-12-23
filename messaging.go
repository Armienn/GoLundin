package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"

	"encoding/json"

	"github.com/Armienn/GoServer"
)

type ThreadData struct {
	MainData
	Thread
}

func NewThreadData(thread Thread, user string, scripts ...string) *ThreadData {
	return &ThreadData{MainData{scripts, user}, thread}
}

type Thread struct {
	ID          int
	Title       string
	MainMessage string
	Responses   []string
}

func loadThreads() []Thread {
	file, err := ioutil.ReadFile("threads.json")
	if err != nil {
		return nil
	}
	var threads []Thread
	err = json.Unmarshal(file, &threads)
	if err != nil {
		return nil
	}
	return threads
}

func threadHandler(server *goserver.Server, w http.ResponseWriter, r *http.Request, path string, session goserver.Session, user interface{}) {
	id, err := strconv.Atoi(path)
	threads := loadThreads()
	var thread Thread
	for _, thread = range threads {
		if thread.ID == id {
			break
		}
	}
	data := NewThreadData(thread, user.(string))
	temp, err := template.ParseFiles("pages/thread.html", "pages/base-start.html", "pages/base-end.html", "pages/header.html")
	if err != nil {
		w.Write([]byte("Fejl: " + err.Error()))
	} else {
		temp.Execute(w, data)
	}
}