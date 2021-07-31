package mylogger

import "fmt"

func ListenForLog(ch chan string) {
	for {
		msg := <-ch
		fmt.Println(msg)
	}
}
