package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.comvtthuan789mygolangcodego-httpclient/gohttp"
)

var (
	githubHttpClient = getGithubHttpClient()
)

func getGithubHttpClient() gohttp.HttpClient {
	client := gohttp.New()

	client.SetMaxIdleConnections(20)
	client.SetConnectionTimeout(2 * time.Second)
	client.SetResponseTimeout(50 * time.Millisecond)

	client.DisableTimeouts(true)

	commonHeaders := make(http.Header)
	client.SetHeaders(commonHeaders)

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

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
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

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}
