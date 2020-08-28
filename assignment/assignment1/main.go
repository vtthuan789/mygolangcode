package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, val := range s {
		if val%2 == 0 {
			fmt.Printf("%d is even\n", val)
		} else {
			fmt.Printf("%d is odd\n", val)
		}
	}
}
