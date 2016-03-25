package main

import (
	"fmt"
)

func main() {
	for n := 2; n < 100; n++ {
		if n % 3 == 0 {
			fmt.Printf("%d ", n)
		}
	}
}
