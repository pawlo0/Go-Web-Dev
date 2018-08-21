package main

import (
	"html/template"
	"log"
	"os"
)

type hotel struct {
	Name, Address, City, Zip string
}

type region struct {
	Region string
	Hotels []hotel
}

type regions []region

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("tpl.gohtml"))
}

func main() {

	hotels := regions{
		region{
			Region: "Southern",
			Hotels: []hotel{
				hotel{"hotel one", "1 high st", "southern city", "1224"},
				hotel{"hotel two", "2 high st", "another southern city", "2264"},
				hotel{"hotel three", "23 corner st", "southern city", "1225"},
			},
		},
		region{
			Region: "Central",
			Hotels: []hotel{
				hotel{"hotel four", "22 navigation st", "a central city", "1211"},
				hotel{"hotel five", "1 low st", "another central city", "24"},
			},
		},
		region{
			Region: "Northern",
			Hotels: []hotel{
				hotel{"hotel six", "1 memory st", "northern city", "2222"},
				hotel{"hotel seven", "44 oak st", "some city", "3333"},
				hotel{"hotel eight", "1 high st", "another northern city", "4444"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
