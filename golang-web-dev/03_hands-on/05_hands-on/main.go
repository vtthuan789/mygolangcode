package main

import (
	"log"
	"net/http"
	"text/template"
)

func foo(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("starting-files/templates/index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	tpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/pics/", http.FileServer(http.Dir("starting-files/public")))
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
