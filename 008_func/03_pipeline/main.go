package main

import (
	"html/template"
	"log"
	"math"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func double(t int) int {
	return t * 2
}

func square(t int) float64 {
	return float64(t * t)
}

func sqroot(t float64) float64 {
	return math.Sqrt(t)
}

var fm = template.FuncMap{
	"double": double,
	"square": square,
	"sqroot": sqroot,
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 3)
	if err != nil {
		log.Fatal(err)
	}
}
