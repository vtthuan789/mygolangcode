package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	// get REQUEST LINE
	scanner.Scan()
	requestLine := scanner.Text()
	s := strings.Fields(requestLine)
	requestMethod := s[0]
	requestURI := s[1]
	fmt.Println("REQUEST method:", requestMethod)
	fmt.Println("REQUEST URI:", requestURI)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		fmt.Println(line)
	}

	var body string

	switch {
	case requestMethod == "GET" && requestURI == "/":
		body = `
			<h1>"GET INDEX"</h1>
			<a href="/">index</a><br>
			<a href="/apply">apply</a><br>
			`
	case requestMethod == "GET" && requestURI == "/apply":
		body = `
			<h1>"GET APPLY"</h1>
			<a href="/">index</a><br>
			<a href="/apply">apply</a><br>
			<form action="/apply" method="POST">
			<input type="hidden" value="In my good death">
			<input type="submit" value="submit">
			</form>
			`
	case requestMethod == "POST" && requestURI == "/apply":
		body = `
			<h1>"POST APPLY"</h1>
			<a href="/">index</a><br>
			<a href="/apply">apply</a><br>
			`
	default:
		body = `<h1>HOLY COW THIS IS LOW LEVEL</h1>`
	}

	handle(conn, body)

}

func handle(conn net.Conn, htmlBody string) {

	body := fmt.Sprintf(`<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Document</title>
	</head>
	<body>
	%s
	</body>
	</html>`, htmlBody)

	_, err := io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))

	fmt.Fprint(conn, "Content-Type: text/html\r\n")

	_, err = io.WriteString(conn, "\r\n")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = io.WriteString(conn, body)
	if err != nil {
		log.Fatalln(err)
	}
}
