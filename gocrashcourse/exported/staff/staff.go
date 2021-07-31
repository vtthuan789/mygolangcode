package staff

import "fmt"

var OverpaidLimit = 50000
var underpaidLimit = 20000

type Employee struct {
	FirstName string
	LastName  string
	Salary    int
	FullTime  bool
}
type Office struct {
	AllStaff []Employee
}

func (e *Office) All() []Employee {
	return e.AllStaff
}

func (e *Office) Overpaid() []Employee {
	var overpaid []Employee

	for _, x := range e.AllStaff {
		if x.Salary >= OverpaidLimit {
			overpaid = append(overpaid, x)
		}
	}
	return overpaid
}

func (e *Office) Underpaid() []Employee {
	var underpaid []Employee

	for _, x := range e.AllStaff {
		if x.Salary < underpaidLimit {
			underpaid = append(underpaid, x)
		}
	}
	return underpaid
}

func (e *Office) notVisible() {
	fmt.Println("Avenger, Assemble!")
}
