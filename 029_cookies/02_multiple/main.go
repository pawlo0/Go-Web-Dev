package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundance", abudance)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
	})

	fmt.Fprintln(w, "Cookie written - check your browser")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

func read(w http.ResponseWriter, req *http.Request) {

	c1, err := req.Cookie("my-cookie")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "Your Cookie #1:", c1)
	}

	c2, err := req.Cookie("general")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "Your Cookie #2:", c2)
	}

	c3, err := req.Cookie("specific")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "Your Cookie #3:", c3)
	}
}

func abudance(w http.ResponseWriter, req *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:  "general",
		Value: "some other value about general things",
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "specific",
		Value: "some other value about specific things",
	})

	fmt.Fprintln(w, "Cookie written - check your browser")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}
