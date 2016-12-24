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

func threadGetHandler(w http.ResponseWriter, r *http.Request, info goserver.Info) {
	id, err := strconv.Atoi(info.Path)
	if err != nil {
		w.Write([]byte("Fejl: " + err.Error()))
		return
	}
	threads := loadThreads()
	var thread Thread
	for _, thread = range threads {
		if thread.ID == id {
			break
		}
	}
	data := NewThreadData(thread, info.User())
	temp, err := template.ParseFiles("pages/thread.html", "pages/base-start.html", "pages/base-end.html", "pages/header.html")
	if err != nil {
		w.Write([]byte("Fejl: " + err.Error()))
	} else {
		temp.Execute(w, data)
	}
}

func threadPostHandler(w http.ResponseWriter, r *http.Request, info goserver.Info) {
	threads := loadThreads()
	nextID := 0
	for _, thread := range threads {
		if nextID <= thread.ID {
			nextID = thread.ID + 1
		}
	}
	thread := Thread{}
	thread.ID = nextID
	var ok bool
	r.ParseForm()
	thread.Title, ok = FromForm(r, "title")
	if !ok {
		w.Write([]byte("Fejl: Mangler titel"))
		return
	}
	thread.MainMessage, ok = FromForm(r, "message")
	if !ok {
		w.Write([]byte("Fejl: Mangler besked"))
		return
	}
	threads = append(threads, thread)
	jsonThreads := toJson(threads)
	ioutil.WriteFile("threads.json", []byte(jsonThreads), 0)
	http.Redirect(w, r, "/beskeder/"+strconv.Itoa(nextID), http.StatusTemporaryRedirect)
}
