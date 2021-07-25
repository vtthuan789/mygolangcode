package mypackage

import "fmt"

var privateVar = "This is my private variable"
var PublicVar = "This is my public variable"
var PackageVar = "This is mypackage variable"

func notExported() {
	println("This is not exported function")
}
func Exported() {
	println("This is exported function")
}
func PrintMe(packageVar, blockLevelVar string) {
	fmt.Println(PackageVar, packageVar, blockLevelVar)
}
