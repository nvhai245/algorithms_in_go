package main

import (
	"fmt"
	"github.com/nvhai245/algorithms_in_go/libs"
	"math/rand"
	"time"
)

func sort(a []int) {
	if len(a) < 2 {
		return
	}

	left, right := 0, len(a)-1
	pivot := rand.Int() % len(a)
	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]
	sort(a[:left])
	sort(a[left+1:])
}

func quickSort(a []int) []int {
	defer libs.TimeTrack(time.Now(), "quickSort")
	sort(a)
	return a
}

func main() {
	fmt.Println(quickSort(libs.GenerateSlice(10)))
}
