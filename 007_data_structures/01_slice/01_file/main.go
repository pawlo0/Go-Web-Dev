package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	sages := []string{"Gandhi", "MLH", "Buddha", "Jesus", "Muhammad"}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("Error creating file", err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
