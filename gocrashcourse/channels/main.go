package main

import (
	"fmt"
	"strings"

	"github.com/eiannone/keyboard"
)

var keyPressChan chan rune

func main() {
	keyPressChan = make(chan rune)

	go listenForKeyPress()

	fmt.Println("Press any key, or q to quit")
	_ = keyboard.Open()
	defer func() {
		keyboard.Close()
	}()

	for {
		char, _, _ := keyboard.GetSingleKey()
		if strings.ToLower(string(char)) == "q" {
			break
		}

		keyPressChan <- char
	}
}

func listenForKeyPress() {
	for {
		key := <-keyPressChan
		fmt.Println("You press", string(key))
	}
}
