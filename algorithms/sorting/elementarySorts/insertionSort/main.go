package main

import (
	"fmt"
	"github.com/nvhai245/algorithms_in_go/libs"
	"time"
)

func insertionSort(a []int) []int {
	defer libs.TimeTrack(time.Now(), "insertionSort")
	for i := 1; i < len(a); i++ {
		for j := i; j > 0 && a[j] < a[j-1]; j-- {
			k := a[j-1]
			a[j-1] = a[j]
			a[j] = k
		}
	}
	return a
}

func main() {
	fmt.Println("[UNSORTED]:")
	fmt.Println(libs.TestIntSlice)
	fmt.Println(insertionSort(libs.TestIntSlice))
}
