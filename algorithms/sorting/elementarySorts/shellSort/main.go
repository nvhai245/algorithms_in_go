package main

import (
	"fmt"
	"github.com/nvhai245/algorithms_in_go/libs"
	"time"
)

func shellSort(a []int) []int {
	defer libs.TimeTrack(time.Now(), "shellSort")
	h := 1
	for h < len(a)/3 {
		h = h*3 + 1
	}
	for h >= 1 {
		for i := h; i < len(a); i++ {
			for j := i; j >= h && a[j] < a[j-h]; j -= h {
				k := a[j-h]
				a[j-h] = a[j]
				a[j] = k
			}
		}
		h = h / 3
	}
	return a
}

func main() {
	fmt.Println(shellSort(libs.GenerateSlice(20)))
}
