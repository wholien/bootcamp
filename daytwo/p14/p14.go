package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scanf("%s", &input)
	// fmt.Printf("%s", input)	
	length := len(input)
	interval := 7
	var j int
	for i := 0; i < interval; i++ {
		j = i
		for j < length {
			fmt.Printf(string(input[j]))
			j += interval
		}
		fmt.Printf("\n")
	}
}
