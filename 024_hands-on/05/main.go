package main

import (
	"log"
	"net/http"

	"github.com/golang/go/src/pkg/html/template"
)

func main() {
	http.HandleFunc("/", renderTemplate)
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/pics/", fs)
	http.ListenAndServe(":8080", nil)
}

func renderTemplate(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("templates/index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
