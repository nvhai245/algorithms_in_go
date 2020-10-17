package main

import (
	"fmt"
	"github.com/nvhai245/algorithms_in_go/libs"
	"time"
)

func selectionSort(a []int) []int {
	defer libs.TimeTrack(time.Now(), "selectionSort")
	for i := 0; i < len(a); i++ {
		min := a[i]
		k := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < min {
				min = a[j]
				k = j
			}
		}
		a[k] = a[i]
		a[i] = min
	}
	return a
}

func main() {
	fmt.Println(selectionSort(libs.GenerateSlice(20)))
}
