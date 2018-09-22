package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", set)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.SetCookie(w, &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
		})
		fmt.Fprintln(w, "This is the first time you visit this site.")
	} else {
		v, _ := strconv.Atoi(c.Value)
		v++
		http.SetCookie(w, &http.Cookie{
			Name:  "my-cookie",
			Value: strconv.Itoa(v),
		})
		fmt.Fprintln(w, "Number of visits to this site:", c.Value)
	}

}
