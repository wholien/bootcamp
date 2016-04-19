package main

import "fmt"

func main() {
	sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
}

func sum(slicey []int) {
	summ := 0;
	for i := 0; i < len(slicey); i++ {
		summ += slicey[i]
	}
	fmt.Println(summ)
}
		
