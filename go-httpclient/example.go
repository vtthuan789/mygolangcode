package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.comvtthuan789mygolangcodego-httpclient/gohttp"
)

var (
	githubHttpClient = getGithubHttpClient()
)

func getGithubHttpClient() gohttp.HttpClient {
	client := gohttp.New()

	commonHeaders := make(http.Header)
	commonHeaders.Set("Authorization", "Bearer ABC 123")
	client.SetHeaders(commonHeaders)

	return client
}

func main() {
	getUrls()
	getUrls()
	getUrls()
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
