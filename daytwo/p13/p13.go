package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	input := os.Args[1]
	desired := os.Args[2]
	if strings.HasSuffix(input, "mi") {
		orig, _ := strconv.Atoi(strings.TrimSuffix(input, "mi"))
		switch desired {
		case "km": fmt.Printf("%.2fkm\n", (float64(orig) * 1.6))
		}
	} else if strings.HasSuffix(input, "km") {
	} else if strings.HasSuffix(input, "m") {
	} else if strings.HasSuffix(input, "ft") {
	}
}
	
