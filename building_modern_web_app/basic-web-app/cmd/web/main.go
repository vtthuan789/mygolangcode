package main

import (
	"fmt"
	"net/http"

	"github.com/vtthuan789/mygolangcode/building_modern_web_app/basic-web-app/pkg/handlers"
)

const portNumber = ":8080"

// main is the main function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
