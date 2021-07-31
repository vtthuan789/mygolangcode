package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	coffees := make(map[int]string)
	coffees[1] = "Cappucino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Espresso"

	fmt.Println("MENU")
	fmt.Println("----")
	fmt.Println("1 - Cappucino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Americano")
	fmt.Println("4 - Mocha")
	fmt.Println("5 - Macchiato")
	fmt.Println("6 - Espresso")
	fmt.Println("Q - Quit the program")

	char := ' '
	for strings.ToLower(string(char)) != "q" {
		char, _, err = keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		i, _ := strconv.Atoi(string(char))
		if val, ok := coffees[i]; ok {
			fmt.Println(fmt.Sprintf("You chose %s", val))
		} else {
			fmt.Println(fmt.Sprintf("You chose %c", char))
		}

	}

	fmt.Println("Program existing...")
}
