package main

import "fmt"

func main() {
	fmt.Printf("How many miles?\n")
	var miles float32
	fmt.Scanf("%f", &miles)
	fmt.Printf("<!DOCTYPE html>\n")
	fmt.Printf("<html>\n")
	fmt.Printf("\t<head></head>\n")
	fmt.Printf("\t<body>\n")
	fmt.Printf("\t\tMiles: %.2f<br>\n", miles)
	fmt.Printf("\t\tKilometers: %.2f\n", miles/1.6)
	fmt.Printf("\t</body>\n")
	fmt.Printf("</html>\n")
}
