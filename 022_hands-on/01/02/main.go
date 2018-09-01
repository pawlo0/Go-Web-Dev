package main

import (
	"io"
	"net/http"
)

func i(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Home page")
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func me(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "My name is Pawlo")
}

func main() {

	http.HandleFunc("/", i)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me", me)
	http.ListenAndServe(":8080", nil)
}
