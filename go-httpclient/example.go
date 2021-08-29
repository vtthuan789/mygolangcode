package main

import (
	"fmt"
	"io/ioutil"

	"github.comvtthuan789mygolangcodego-httpclient/gohttp"
)

func main() {
	client := gohttp.New()

	response, err := client.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}
