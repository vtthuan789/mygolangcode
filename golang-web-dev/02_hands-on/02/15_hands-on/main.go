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
	fmt.Println("REQUEST method:", s[0])
	fmt.Println("REQUEST URI:", s[1])

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		fmt.Println(line)
	}

	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Document</title>
	</head>
	<body>
	
	<h1>HOLY COW THIS IS LOW LEVEL</h1>
	
	</body>
	</html>`

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
