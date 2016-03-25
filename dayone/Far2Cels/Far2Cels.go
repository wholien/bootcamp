package main

import (
	"fmt"
)

func main() {
	fmt.Printf("What degress Farhenheit?\n")
	var degrees float32
	fmt.Scanf("%f", &degrees)
	fmt.Printf("That is %f degrees Celsius\n", (degrees-32)*5/9)
}
