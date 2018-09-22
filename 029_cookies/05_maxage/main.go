package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, `<a href="/set">Set a cookie</a>`)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
	})
	fmt.Fprintln(w, `<a href="/read">read</a>`)
}

func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Redirect(w, req, "/set", http.StatusSeeOther)
		return
	}
	fmt.Fprintf(w, `<h1>Your Cookie:<br>%v</h1><h1><a href="/expire">expire</a></h1>`, c)
}

func expire(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Redirect(w, req, "/set", http.StatusSeeOther)
		return
	}
	c.MaxAge = -1
	http.SetCookie(w, c)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}
