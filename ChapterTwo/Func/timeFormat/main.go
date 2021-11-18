package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

var ft = template.FuncMap{
	"tf": tformat,
}

func init() {
	tpl = template.Must(template.New("").Funcs(ft).ParseFiles("templates/tpl.gohtml"))
}

func tformat(a time.Time) string {
	return a.Format("01/02 2006 03:04:05")
}

func main() {
	z := time.Now()
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
