package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "dog.gohtml", nil)
}

func dogPic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dog.jpeg")
}

func main() {
	http.Handle("/", http.HandlerFunc(foo))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/dog.jpeg", http.HandlerFunc(dogPic))

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
