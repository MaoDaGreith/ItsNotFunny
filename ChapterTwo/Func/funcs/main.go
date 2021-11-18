package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

var ft = template.FuncMap{
	"db": double,
}

func init() {
	tpl = template.Must(template.New("").Funcs(ft).ParseFiles("templates/tpl.gohtml"))
}

func double(a int) int {
	return a * a
}

func main() {
	z := []int{1, 2, 3, 4, 5}

	fl, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer fl.Close()

	err = tpl.ExecuteTemplate(fl, "tpl.gohtml", z)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", z)
	if err != nil {
		log.Fatalln(err)
	}
}
