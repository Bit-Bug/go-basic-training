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
	for j := 0; j <= length-1; j++ {
		var iMin = j
		for i := j + 1; i < length; i++ {
			if arrays[i] < arrays[iMin] {
				iMin = i
			}

		}
		if iMin != j {
			var max = arrays[iMin]
			arrays[iMin] = arrays[j]
			arrays[j] = max
		}

	}
}
