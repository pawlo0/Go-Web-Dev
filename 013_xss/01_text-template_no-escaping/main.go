package main

import (
	"log"
	"os"
	"text/template"
)

type Page struct {
	Title, Heading, Input string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("tpl.gohtml"))
}

func main() {

	home := Page{
		Title:   "Nothing escaped",
		Heading: "Nothing is escaped with text/template",
		Input:   `<script>alert("Wow")</script>`,
	}

	err := tpl.Execute(os.Stdout, home)
	if err != nil {
		log.Fatalln(err)
	}
}
