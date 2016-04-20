package main

import "fmt"

func centeredAverage(intlist []float64) float64 {
	length := len(intlist)
	sm := intlist[0]
	lg := intlist[0]
	var sum float64
	for _, k := range intlist {
		if k < sm {sm = k}
		if k > lg {lg = k}
		sum += k
	}
	sum -= (lg + sm)
//	fmt.Println(sum)
	return sum/float64(length-2)
}

func main() {
	avg := centeredAverage([]float64{1, 2, 3, 4, 100})
	fmt.Println(avg)
}
