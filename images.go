package main

import (
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/Armienn/GoServer"
)

func imagesPostHandler(w http.ResponseWriter, r *http.Request, info goserver.Info) {
	err := r.ParseMultipartForm((1 << 20) * 100) //max 100 megabyte
	if err != nil {
		w.Write([]byte("Fejl: " + err.Error()))
		return
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
			if outfile, err = os.Create("./uploaded/" + fileHeader.Filename); nil != err {
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

func imagesGetHandler(w http.ResponseWriter, r *http.Request, info goserver.Info) {
	data := NewMainData(info.User())
	temp, err := template.ParseFiles("pages/images.html", "pages/base-start.html", "pages/base-end.html", "pages/header.html", "pages/sidebar.html")
	if err != nil {
		w.Write([]byte("Fejl: " + err.Error()))
	} else {
		temp.Execute(w, data)
	}
}
