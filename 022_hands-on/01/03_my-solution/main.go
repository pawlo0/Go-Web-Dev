package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {

	http.HandleFunc("/", i)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me", me)
	http.ListenAndServe(":8080", nil)
}

func i(res http.ResponseWriter, req *http.Request) {
	data := "This is the home page"
	tpl.ExecuteTemplate(res, "index.gohtml", data)
}

func dog(res http.ResponseWriter, req *http.Request) {
	data := "This is the dog page AUWF AUWF"
	tpl.ExecuteTemplate(res, "index.gohtml", data)
}

func me(res http.ResponseWriter, req *http.Request) {
	data := "My name is Paulo"
	tpl.ExecuteTemplate(res, "index.gohtml", data)
}
