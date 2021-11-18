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

type numeros struct {
	Pjerdoles int
	Segnores  string
}

type gustavos struct {
	Exam   []exampl
	Numero []numeros
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	z := exampl{
		Name:    "jenius",
		Zaebist: "Hovna",
	}
	x := exampl{
		Name:    "Hesus",
		Zaebist: "Derimisha",
	}
	c := exampl{
		Name:    "Dang",
		Zaebist: "God",
	}

	a := numeros{
		Pjerdoles: 2,
		Segnores:  "Urangutanos",
	}

	s := numeros{
		Pjerdoles: 5,
		Segnores:  "Snakeos",
	}

	d := numeros{
		Pjerdoles: 5,
		Segnores:  "Snakeos",
	}

	v := []exampl{z, x, c}
	f := []numeros{a, s, d}

	data := gustavos{
		Exam:   v,
		Numero: f,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}

	fl, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer fl.Close()

	err = tpl.Execute(fl, data)
	if err != nil {
		log.Fatalln(err)
	}
}
