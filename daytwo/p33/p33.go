package main

import "fmt"

func main() {
	fmt.Println(greatest(1, 2, 3, 4))
}

func greatest(nums ...int) int {
	biggest := nums[0]
	for _, num := range nums {
		if num > biggest {
			biggest = num
		}
	}
	return biggest
}
