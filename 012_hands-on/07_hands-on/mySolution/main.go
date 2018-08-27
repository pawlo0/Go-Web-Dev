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

type restaurant struct {
	Name     string
	Location string
	Menu     menu
}

type restaurants []restaurant

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("tpl.gohtml"))
}

func main() {

	r := restaurants{
		restaurant{
			Name:     "First restaurant",
			Location: "1 High St",
			Menu: menu{
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
			},
		},
		restaurant{
			Name:     "Second restaurant",
			Location: "23 Corner St",
			Menu: menu{
				meal{
					Meal: "Breakfast",
					Item: []item{
						item{"Oatmeal"},
						item{"Cheerios"},
						item{"Juice Orange"},
					},
				},
				meal{
					Meal: "Lunch",
					Item: []item{
						item{"Hamburger"},
						item{"Cheese Melted Sandwich"},
						item{"French Fries"},
					},
				},
				meal{
					Meal: "Dinner",
					Item: []item{
						item{"Pasta Bolognese"},
						item{"Steak"},
						item{"Bistro Potatoe"},
					},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, r)
	if err != nil {
		log.Fatalln(err)
	}
}
