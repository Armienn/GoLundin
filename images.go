package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	"strings"

	"github.com/Armienn/GoServer"
)

func imagesPostHandler(w http.ResponseWriter, r *http.Request, info goserver.Info) {
	err := r.ParseMultipartForm((1 << 20) * 100) //max 100 megabyte
	if err != nil {
		w.Write([]byte("Fejl: " + err.Error()))
		return
	}
	path := "files/images/" + info.User() + "/"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}

	for _, fileHeaders := range r.MultipartForm.File {
		for _, fileHeader := range fileHeaders {
			var file multipart.File
			if file, err = fileHeader.Open(); err != nil {
				w.Write([]byte("Fejl: " + err.Error()))
				return
			}
			defer file.Close()
			var outfile *os.File
			if outfile, err = os.Create(path + fileHeader.Filename); nil != err {
				w.Write([]byte("Fejl: " + err.Error()))
				return
			}
			defer outfile.Close()
			if _, err = io.Copy(outfile, file); nil != err {
				w.Write([]byte("Fejl: " + err.Error()))
				return
			}
		}
	}
	imagesGetHandler(w, r, info)
}

type FileData struct {
	MainData
	Images      []string
	Directories []string
}

func NewFileData(directory string, subDirectory string, user string, scripts ...string) *FileData {
	files, _ := ioutil.ReadDir(directory)
	images := []string{}
	directories := []string{}
	for _, file := range files {
		if file.IsDir() {
			directories = append(directories, file.Name())
		}
	}
	files, _ = ioutil.ReadDir(directory + subDirectory)
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), "jpg") || strings.HasSuffix(file.Name(), "png") {
			if subDirectory == "" {
				images = append(images, file.Name())
			} else {
				images = append(images, subDirectory+"/"+file.Name())
			}
		}
	}
	return &FileData{MainData{scripts, user}, images, directories}
}

func imagesGetHandler(w http.ResponseWriter, r *http.Request, info goserver.Info) {
	if info.Path == "ny" {
		showNewImagePage(w, info)
		return
	}
	data := NewFileData("files/images/", info.Path, info.User())
	temp, err := template.ParseFiles("pages/images.html", "pages/base-start.html", "pages/base-end.html", "pages/header.html", "pages/sidebar-images.html")
	if err != nil {
		w.Write([]byte("Fejl: " + err.Error()))
	} else {
		temp.Execute(w, data)
	}
}

func showNewImagePage(w http.ResponseWriter, info goserver.Info) {
	data := NewFileData("files/images/", "", info.User())
	temp, err := template.ParseFiles("pages/images-new.html", "pages/base-start.html", "pages/base-end.html", "pages/header.html", "pages/sidebar-images.html")
	if err != nil {
		w.Write([]byte("Fejl: " + err.Error()))
	} else {
		temp.Execute(w, data)
	}
}
