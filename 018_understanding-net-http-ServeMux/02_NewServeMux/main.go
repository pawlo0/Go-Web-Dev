package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

type hotcat int

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

type index int

func (i index) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "home page")
}

func main() {
	var d hotdog
	var c hotcat
	var a index

	mux := http.NewServeMux()
	mux.Handle("/", a)
	mux.Handle("/dog/", d)
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux)
}
