package main

import "fmt"

func reverse(arr []int) []int {
	ys := make([]int, len(arr))
	for k := range arr {
		ys[len(arr) - k - 1] = arr[k] 
	}
	return ys
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	arr2 := reverse(arr)
	fmt.Println(arr2)
}
