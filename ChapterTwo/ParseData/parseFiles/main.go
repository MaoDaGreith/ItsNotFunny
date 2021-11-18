package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("one.gof")
	if err != nil {
		log.Fatalln(err)
	}

	fl, err := os.Create("tpl.gof")
	if err != nil {
		log.Fatalln(err)
	}
	defer fl.Close()

	tpl, err = tpl.ParseFiles("two.gof", "three.gof")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gof", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "three.gof", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gof", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(fl, "three.gof", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
