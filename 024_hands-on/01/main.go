package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "Foo ran")
}

func dog(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}

	data := map[string]string{
		"title": "This is from dog",
		"image": "/dog.jpg",
	}

	err = tpl.ExecuteTemplate(w, "dog.gohtml", data)
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "dog.jpg")
}
