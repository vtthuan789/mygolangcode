package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	californiaHotels := []struct {
		Name    string
		Address string
		City    string
		Zip     string
		Region  string
	}{
		{
			Name:    "Disney's Grand Californian Hotel & Spa",
			Address: "1600 Disneyland Dr",
			City:    "Anaheim",
			Zip:     "CA 92802",
			Region:  "Southern",
		},
		{
			Name:    "Disneyland Hotel",
			Address: "1150 Magic Way",
			City:    "Anaheim",
			Zip:     "CA 92802",
			Region:  "Southern",
		},
		{
			Name:    "Freehand Los Angeles",
			Address: "416 W 8th St",
			City:    "Los Angeles",
			Zip:     "CA 90014",
			Region:  "Northern",
		},
	}

	err := tpl.Execute(os.Stdout, californiaHotels)
	if err != nil {
		log.Fatalln(err)
	}
}
