package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}
type japanBot struct{}

func main() {
	var eb englishBot
	var sb spanishBot
	var jb japanBot

	printGreeting(eb)
	printGreeting(sb)
	printGreeting(jb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func (japanBot) getGreeting() string {
	return "Kon'nichiwa!"
}
