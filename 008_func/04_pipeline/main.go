package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func dayMonthYear(t time.Time) string {
	date := fmt.Sprintln(t.Weekday(), t.Format("02-01-2006"))
	return date
}

var fm = template.FuncMap{
	"ddMMYYY": dayMonthYear,
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatal(err)
	}
}
