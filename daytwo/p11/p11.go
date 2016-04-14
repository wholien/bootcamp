package main

import "fmt"

func main() {
	fmt.Printf("How many Miles?\n")
	var miles float32
	fmt.Scanf("%f", &miles)
	divider := "+--------------------+"
	fmt.Printf("%s\n", divider)
	fmt.Printf("| Miles: %.2f       |\n", miles)
	fmt.Printf("%s\n", divider)
	fmt.Printf("| Kilometers: %.2f   |\n", miles/1.6)
	fmt.Printf("%s\n", divider)
}
