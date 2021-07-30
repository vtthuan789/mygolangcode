package main

import "fmt"

type Vehicle struct {
	NumberOfWheels     int
	NumberOfPassengers int
}

type Car struct {
	Make       string
	Model      string
	Year       int
	IsElectric bool
	IsHybrid   bool
	Vehicle    Vehicle
}

func (v Vehicle) showDetails() {
	fmt.Println("Number of passengers:", v.NumberOfPassengers)
	fmt.Println("Number of wheels:", v.NumberOfWheels)
}

func (c Car) show() {
	fmt.Println("Make:", c.Make)
	fmt.Println("Model:", c.Model)
	fmt.Println("Year:", c.Year)
	fmt.Println("Is Electric:", c.IsElectric)
	fmt.Println("Is Hybrid:", c.IsHybrid)
	c.Vehicle.showDetails()
}

func main() {
	suv := Vehicle{
		NumberOfPassengers: 7,
		NumberOfWheels:     4,
	}
	sorento := Car{
		Make:       "Kia",
		Model:      "Sorento",
		Year:       2021,
		IsElectric: false,
		IsHybrid:   false,
		Vehicle:    suv,
	}

	sorento.show()

	fmt.Println()

	teslaModelX := Car{
		Make:       "Tesla",
		Model:      "Model X",
		Year:       2021,
		IsElectric: true,
		IsHybrid:   false,
		Vehicle:    suv,
	}

	teslaModelX.Vehicle.NumberOfPassengers = 6

	teslaModelX.show()

}
