package main

import (
	"log"
	"os"
	"text/template"
)

type hotels struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
}

type region struct {
	Region string
	Hotels []hotels
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	region := []region{
		{
			Region: "Southern",
			Hotels: []hotels{
				{
					Name:    "MahDik",
					Address: "is hard 69",
					City:    "Sexville",
					Zip:     "69420",
					Region:  "southern",
				}, {
					Name:    "Mapoosi",
					Address: "is wet 69",
					City:    "weeberhood",
					Zip:     "200101",
					Region:  "southern",
				},
			},
		}, {
			Region: "Northern",
			Hotels: []hotels{
				{
					Name:    "Maback",
					Address: "is hurt 420",
					City:    "Manhood",
					Zip:     "42069",
					Region:  "northern",
				}, {
					Name:    "Mahbouner",
					Address: "is long 69",
					City:    "Gayrespect",
					Zip:     "22422",
					Region:  "northern",
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, region)
	if err != nil {
		log.Fatalln(err)
	}
}
