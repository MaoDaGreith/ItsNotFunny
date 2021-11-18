package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	z := []int{1, 2, 3, 4, 5}
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
