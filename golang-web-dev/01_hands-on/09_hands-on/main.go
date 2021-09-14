package main

import (
	"encoding/csv"
	"log"
	"os"
	"text/template"
)

var functions = template.FuncMap{
	"iterate": Iterate,
}

func Iterate(count int) []int {
	var i int
	var items []int
	for i = 0; i < count; i++ {
		items = append(items, i)
	}
	return items
}

type MenuItem struct {
	Name        string
	Price       float32
	Description string
}

type Menu struct {
	Meal  string
	Items []MenuItem
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(functions).ParseFiles("tpl.gohtml"))
}

func main() {
	f, err := os.Open("./table.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}
	recordsMap := make(map[string][]string)
	keys := records[0]

	for _, record := range records[1:] {
		for index, value := range record[:2] {
			recordsMap[keys[index]] = append(recordsMap[keys[index]], value)
		}
	}

	htmlFile, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer htmlFile.Close()

	err = tpl.ExecuteTemplate(htmlFile, "tpl.gohtml", recordsMap)
	if err != nil {
		log.Fatalln(err)
	}
}
