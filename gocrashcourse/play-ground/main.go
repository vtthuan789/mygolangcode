package main

import (
	"log"

	"github.com/vtthuan789/mygolangcode/gocrashcourse/helpers"
)

func main() {
	var myType helpers.SomeType
	myType.TypeString = "Some string"
	myType.TypeInt = 79
	log.Println(myType)
	// you can't do this in golang
	// var myString * string
	// myString  = "Read"
	var myString = "Read"
	log.Println("Before:", myString)
	log.Println("Adress Before:", &myString)
	changeString(&myString)
	// myString = "Write"
	log.Println("After:", myString)
	log.Println("Adress After:", &myString)
}

func changeString(s *string) {
	*s = "Write"
}
