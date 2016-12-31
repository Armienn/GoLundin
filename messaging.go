package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"encoding/json"

	"github.com/Armienn/GoServer"
	"github.com/russross/blackfriday"
)

type ForumData struct {
	MainData
	Threads []Thread
}

func NewForumData(threads []Thread, user string, scripts ...string) *ForumData {
	return &ForumData{MainData{scripts, user}, threads}
}

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
	Section     string
	Responses   []Thread
	Author      string
	Time        time.Time
}

func (thread Thread) Markdown() template.HTML {
	return template.HTML(string(blackfriday.MarkdownBasic([]byte(thread.MainMessage))))
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
	if info.Path == "ny" {
		showNewThreadPage(w, info)
		return
	}
	id, err := strconv.Atoi(info.Path)
	if err != nil {
		showSection(w, info)
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
	temp, err := template.ParseFiles("pages/thread.html", "pages/base-start.html", "pages/base-end.html", "pages/header.html", "pages/sidebar.html")
	if err != nil {
		w.Write([]byte("Fejl: " + err.Error()))
	} else {
		temp.Execute(w, data)
	}
}

func showSection(w http.ResponseWriter, info goserver.Info) {
	threads := loadThreads()
	section := []Thread{}
	for _, thread := range threads {
		if thread.Section == info.Path || (info.Path == "andet" && thread.Section == "") {
			section = append(section, thread)
		}
	}
	data := NewForumData(section, info.User())
	temp, err := template.ParseFiles("pages/frontpage.html", "pages/base-start.html", "pages/base-end.html", "pages/header.html", "pages/sidebar.html")
	if err != nil {
		w.Write([]byte("Fejl: " + err.Error()))
	} else {
		temp.Execute(w, data)
	}
}

func showNewThreadPage(w http.ResponseWriter, info goserver.Info) {
	data := NewMainData(info.User())
	temp, err := template.ParseFiles("pages/thread-new.html", "pages/base-start.html", "pages/base-end.html", "pages/header.html", "pages/sidebar.html")
	if err != nil {
		w.Write([]byte("Fejl: " + err.Error()))
	} else {
		temp.Execute(w, data)
	}
}

func threadPostHandler(w http.ResponseWriter, r *http.Request, info goserver.Info) {
	r.ParseForm()
	thread, ok := ThreadFromForm(w, r, info)
	if !ok {
		return
	}
	threads := loadThreads()
	if thread.Title == "response" {
		for i, existingThread := range threads {
			if existingThread.ID == thread.ID {
				if existingThread.Responses == nil {
					existingThread.Responses = make([]Thread, 1)
					existingThread.Responses[0] = *thread
				} else {
					existingThread.Responses = append(existingThread.Responses, *thread)
				}
				threads[i] = existingThread
				break
			}
		}
	} else {
		if thread.ID == 0 {
			//add
			for _, existingThread := range threads {
				if thread.ID <= existingThread.ID {
					thread.ID = existingThread.ID + 1
				}
			}
			threads = append(threads, *thread)
		} else {
			//update
			for i, existingThread := range threads {
				if thread.ID == existingThread.ID {
					threads[i] = *thread
					break
				}
			}
		}
	}
	jsonThreads := toJson(threads)
	ioutil.WriteFile("threads.json", []byte(jsonThreads), 0)
	http.Redirect(w, r, "/beskeder/"+strconv.Itoa(thread.ID), http.StatusTemporaryRedirect)
}

func ThreadFromForm(w http.ResponseWriter, r *http.Request, info goserver.Info) (*Thread, bool) {
	thread := &Thread{}
	id, ok := FromForm(r, "ID")
	if !ok {
		thread.ID = 0
	}
	var err error
	thread.ID, err = strconv.Atoi(id)
	if err != nil {
		thread.ID = 0
	}
	thread.Title, ok = FromForm(r, "title")
	if !ok {
		w.Write([]byte("Fejl: Mangler titel"))
		return nil, false
	}
	thread.MainMessage, ok = FromForm(r, "message")
	if !ok {
		w.Write([]byte("Fejl: Mangler besked"))
		return nil, false
	}
	thread.Section, ok = FromForm(r, "section")
	thread.Author = info.User()
	thread.Time = time.Now()
	return thread, true
}
