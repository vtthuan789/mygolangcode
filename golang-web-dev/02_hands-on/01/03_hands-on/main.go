package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func home(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, "This is my home page")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}
func dog(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, "This is my dog page")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}
func me(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, "My name is Tony")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
