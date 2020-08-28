package main

import (
	"fmt"
	"math"
)

type triangle struct{
	height float64
	base float64
}
type square struct{
	sideLength float64
}

type shape interface{
	getArea() float64
}

func main() {
	tr := triangle{12, 34}
	sq := square{56}

	printArea(tr)
	printArea(sq)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return math.Pow(s.sideLength, 2) 
}