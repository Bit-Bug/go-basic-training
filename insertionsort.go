package main

import (
	"fmt"
)

func main() {
	var scores = []int{90, 70, 50, 80, 60, 85}
	var length = len(scores)
	sort(scores, length)
	for i := 0; i < length; i++ {
		fmt.Printf("%d ", scores[i])
	}

}
func sort(arrays []int, length int) {
	for i := 0; i <= length-1; i++ {
		for j := i; j > 0; j-- {
			if arrays[j-1] > arrays[j] {
				var flag = arrays[j-1]
				arrays[j-1] = arrays[j]
				arrays[j] = flag
			}
		}
	}

}
