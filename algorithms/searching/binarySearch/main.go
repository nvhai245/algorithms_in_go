package main

import (
	"fmt"
	"github.com/nvhai245/algorithms_in_go/libs"
	"time"
)

func binarySearch(needle int, haystack []int) bool {
	defer libs.TimeTrack(time.Now(), "binarySearch")
	low := 0
	high := len(haystack) - 1
	for low <= high {
		median := (low + high) / 2
		if haystack[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	if low == len(haystack) || haystack[low] != needle {
		return false
	}
	return true
}
func main() {
	fmt.Println(binarySearch(63, libs.GenerateSortedSlice(20)))
}
