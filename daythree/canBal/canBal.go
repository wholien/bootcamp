package main

import "fmt"

func sum(sli []int) int {
	sum := 0
	for _, num := range sli {
		sum += num
	}
	return sum
}

func canBal(arr []int) bool {
	for index, _ := range arr {
		if sum(arr[:index]) == sum(arr[index:]) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(canBal([]int{1, 1, 1, 2, 1}))
	fmt.Println(canBal([]int{2,1,1,2,1}))
	fmt.Println(canBal([]int{10,10}))	
}
