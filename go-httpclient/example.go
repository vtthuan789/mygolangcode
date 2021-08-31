package main

import (
	"fmt"

	"github.comvtthuan789mygolangcodego-httpclient/gohttp"
)

var (
	githubHttpClient = getGithubHttpClient()
)

func getGithubHttpClient() gohttp.Client {
	client := gohttp.NewBuilder().
		DisableTimeouts(true).
		SetMaxIdleConnections(5).
		Build()

	return client
}

func main() {
	getUrls()
	createUser(User{
		FirstName: "Thuan",
		LastName:  "Vo",
	})
}

func getUrls() {
	response, err := githubHttpClient.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Status())
	fmt.Println(response.StatusCode())
	fmt.Println(response.String())
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func createUser(user User) {
	response, err := githubHttpClient.Post("https://api.github.com", nil, user)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Status())
	fmt.Println(response.StatusCode())
	fmt.Println(response.Bytes())
}
