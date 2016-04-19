package main

import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fib(n - 1 ) + fib(n - 2)
}

func main() {
	fmt.Printf("%d\n", fib(0))
	fmt.Printf("%d\n", fib(1))
	fmt.Printf("%d\n", fib(2))
	fmt.Printf("%d\n", fib(3))
	fmt.Printf("%d\n", fib(4))
	fmt.Printf("%d\n", fib(5))
}
