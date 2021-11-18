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
	goGovnoTSlucsheVseh := struct {
		Num1 int
		Num2 int
	}{
		Num1: 12,
		Num2: 13,
	}
	err := tpl.Execute(os.Stdout, goGovnoTSlucsheVseh)
	if err != nil {
		log.Fatalln(err)
	}

	fl, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer fl.Close()

	err = tpl.Execute(fl, goGovnoTSlucsheVseh)
	if err != nil {
		log.Fatalln(err)
	}
}
