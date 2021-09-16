package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is my home page")
}
func dog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is my dog page")
}
func me(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "My name is Tony")
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
