package main

import (
	"fmt"

	"myapp/staff"
)

var employees = []staff.Employee{
	{FirstName: "Tony", LastName: "Stark", Salary: 1, FullTime: true},
	{FirstName: "Steve", LastName: "Rogers", Salary: 50000, FullTime: true},
	{FirstName: "Hulk", LastName: "Green", Salary: 70000, FullTime: true},
	{FirstName: "Thor", LastName: "Thunder", Salary: 10000, FullTime: false},
	{FirstName: "Spider", LastName: "Man", Salary: 10000, FullTime: false},
}

func main() {
	myStaff := staff.Office{
		AllStaff: employees,
	}

	staff.OverpaidLimit = 60000

	fmt.Println(myStaff.All())
	fmt.Println("Overpaid Staff", myStaff.Overpaid())
	fmt.Println("Underpaid Staff:", myStaff.Underpaid())
}
