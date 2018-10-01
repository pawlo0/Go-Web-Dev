package main

import (
	"net/http"

	"github.com/golang/go/src/pkg/html/template"
	"github.com/satori/go.uuid"
)

type user struct {
	First string
	Last  string
}

var tpl *template.Template
var dbUsers = map[string]user{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	// get cookie or create in case there isn't a cookie
	cookie, err := r.Cookie("session")
	if err != nil {
		id := uuid.Must(uuid.NewV4())
		cookie := &http.Cookie{
			Name:  "session",
			Value: id.String(),
			// Secure: true,
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, cookie)
	}

	// process form submission
	var u user
	if r.Method == http.MethodPost {
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")
		u = user{f, l}
		dbUsers[cookie.Value] = u
	}

	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, r *http.Request) {

	//get the cookie
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	u, ok := dbUsers[cookie.Value]
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}
