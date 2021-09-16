package main

import (
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

		_, err = io.WriteString(conn, "I see you connected.")
		if err != nil {
			log.Fatalln(err)
		}

		conn.Close()
	}
}
