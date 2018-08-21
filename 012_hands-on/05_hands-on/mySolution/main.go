package main

import (
	"html/template"
	"log"
	"os"
)

type item struct {
	Desc string
}

type meal struct {
	Meal string
	Item []item
}

type menu []meal

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("tpl.gohtml"))
}

func main() {

	m := menu{
		meal{
			Meal: "Breakfast",
			Item: []item{
				item{"Baked beans"},
				item{"Fried eggs"},
			},
		},
		meal{
			Meal: "Lunch",
			Item: []item{
				item{"Pizza"},
				item{"Hamburger"},
			},
		},
		meal{
			Meal: "Dinner",
			Item: []item{
				item{"French Fries"},
				item{"Vegtable soup"},
				item{"pasta bolognese"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, m)
	if err != nil {
		log.Fatalln(err)
	}
}
