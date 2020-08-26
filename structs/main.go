package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	tony := person{
		firstName: "Tony",
		lastName:  "Stark",
		contactInfo: contactInfo{
			email:   "tonystark@stark.com",
			zipCode: 94000,
		},
	}

	tonyPointer := &tony
	tonyPointer.updateName("ironman")
	tony.print()
	fmt.Println(*&tony)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
