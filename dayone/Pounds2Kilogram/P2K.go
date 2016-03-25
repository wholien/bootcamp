package main

import (
	"fmt"
)

func main() {
	fmt.Printf("How many pounds?")
	var pounds float32
	fmt.Scanf("%f", &pounds)
	fmt.Printf("That is %.2f kilos\n", pounds/2.2)
}
