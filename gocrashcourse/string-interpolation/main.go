package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

var reader *bufio.Reader

type User struct {
	UserName        string
	Age             int
	FavouriteNumber float64
	OwnADog         bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	var user User
	user.UserName = readString("What's your name?")
	user.Age = readInt("How old are you?")
	user.FavouriteNumber = readFloat("What is your favourite number?")
	user.OwnADog = readBoolean("Do you own a dog (y/n)?")

	// fmt.Println("Your name is "+userName+". You are", age, "years old.")

	// fmt.Println(fmt.Sprintf("Your name is %s. You are %d years old.", userName, age))
	fmt.Printf("Your name is %s. You are %d years old and your favourite number is %.2f. Owns a dog: %t\n",
		user.UserName,
		user.Age,
		user.FavouriteNumber,
		user.OwnADog)
}

func prompt() {
	fmt.Print("-> ")
}

func readString(s string) string {
	for {
		fmt.Println(s)
		fmt.Print("-> ")

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		if userInput == "" {
			fmt.Println("Please enter a value")
		} else {
			return userInput
		}
	}
}

func readInt(s string) int {
	for {
		fmt.Println(s)
		fmt.Print("-> ")

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)

		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please enter a whole number")
		} else {
			return num
		}
	}
}

func readFloat(s string) float64 {
	for {
		fmt.Println(s)
		fmt.Print("-> ")

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)

		num, err := strconv.ParseFloat(userInput, 64)
		if err != nil {
			fmt.Println("Please enter a number")
		} else {
			return num
		}
	}
}

func readBoolean(s string) bool {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()
	for {
		fmt.Println(s)
		fmt.Print("-> ")
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		s := strings.ToLower(string(char))
		if s == "y" {
			return true
		} else if s == "n" {
			return false
		} else {
			fmt.Println("Please enter y(yes) or n(no)")
		}
	}
}
