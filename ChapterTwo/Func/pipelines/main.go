package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tpl *template.Template

var ft = template.FuncMap{
	"db":   double,
	"sq":   square,
	"sqrt": math.Sqrt,
}

func init() {
	tpl = template.Must(template.New("").Funcs(ft).ParseFiles("templates/tpl.gohtml"))
}

func double(a float64) float64 {
	return a * a
}

func square(a float64) float64 {
	return a * a * a
}

func main() {
	var z float64 = 4

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
