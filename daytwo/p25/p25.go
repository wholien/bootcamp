package main

import (
	"fmt"
	"os"
)

func main() {
	code := os.Args[1]
	m := make(map[string]string)
	m["CA"] = "California"
	fmt.Printf("%s\n", m[code])
}
