package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("asdf asdf asdf")
	fmt.Printf("%v\n", WordCount("asdf asdf asdf\n"))
}

func WordCount(s string) map[string]int {
	strlist := strings.Fields(s)
	strmap := make(map[string]int)
	for _, word := range strlist {
		_, ok := strmap[word]
		if !ok {
			strmap[word] = 1
		} else {
			strmap[word]++
		}

	fmt.Printf("%d", strmap[word])
	}
	return strmap
}	
