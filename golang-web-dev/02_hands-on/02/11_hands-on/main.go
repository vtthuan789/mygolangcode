package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
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
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		fmt.Println(line)
	}

	fmt.Println("Code got here.")
	body := "I see you connected."

	_, err := io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))

	fmt.Fprint(conn, "Content-Type: text/plain\r\n")

	_, err = io.WriteString(conn, "\r\n")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = io.WriteString(conn, body)
	if err != nil {
		log.Fatalln(err)
	}

}
