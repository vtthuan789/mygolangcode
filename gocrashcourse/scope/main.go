package main

import (
	"fmt"
	"myapp/mypackage"
)

var one = "One"
var myVar = "my package level variable"

func main() {
	var one = "This is a block level variable"
	var blockVar = "This is my block level variable"

	fmt.Println(one)

	myFunc()

	fmt.Println(mypackage.PublicVar)

	mypackage.Exported()

	mypackage.PrintMe(myVar, blockVar)
}

func myFunc() {
	fmt.Println(one)
}
