package main

import (
	"fmt"
)

func main() {
	for {
		fmt.Printf("How many miles?\n")
		var miles float32
		fmt.Scanf("%f", &miles)
		fmt.Printf("That is %f kilometers\n", miles*1.60934)
	}
}
