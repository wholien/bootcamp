package main

import (
	"fmt"
)

func main() {
	for n := 1; n <= 100; n++ {
		if n % 3 == 0 {
			fmt.Printf("Fizz")
		}
		if n % 5 == 0 {
			fmt.Printf("Buzz")
		}
		fmt.Printf("\n")
	}
}
