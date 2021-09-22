package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", home)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err == http.ErrNoCookie {
		c = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
			Path:  "/",
		}
	}

	visitedTimes, err := strconv.Atoi(c.Value)
	if err != nil {
		log.Fatalln(err)
	}

	visitedTimes++
	c.Value = strconv.Itoa(visitedTimes)
	http.SetCookie(w, c)

	io.WriteString(w, fmt.Sprintf("You visited here %d times", visitedTimes))
}

// Using cookies, track how many times a user has been to your website domain.
