package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "o oo ooo oooo ooooo oooooo ooooooo"
	fmt.Println(str)
	newStr := replaceNth(str, "oo", "FF", 17)
	fmt.Println(newStr)
}

func replaceNth(s, old, new string, n int) string {
	i := 0
	var searchS string
	for j := 1; j <= n; j++ {
		searchS = s[i:]
		index := strings.Index(searchS, old)
		if index < 0 {
			break
		}
		if j == n {
			s = s[:i] + searchS[:index] + new + searchS[index+len(old):]
		} else {
			i += index + 1
			if i >= len(s) {
				break
			}
		}
	}
	return s
}
