package main

import (
	"log"
	"os"
	"text/template"
)

type exampl struct {
	Name    string
	Zaebist string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	z := []exampl{
		{
			Name:    "jenius",
			Zaebist: "Hovna",
		},
		{
			Name:    "Hesus",
			Zaebist: "Derimisha",
		},
		{
			Name:    "Dang",
			Zaebist: "God",
		},
	}

	err := tpl.Execute(os.Stdout, z)
	if err != nil {
		log.Fatalln(err)
	}

	fl, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer fl.Close()

	err = tpl.Execute(fl, z)
	if err != nil {
		log.Fatalln(err)
	}

}
