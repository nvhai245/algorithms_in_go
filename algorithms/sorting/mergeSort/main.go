package main

import (
	"fmt"
	"github.com/nvhai245/algorithms_in_go/libs"
	"time"
)

func merge(a []int) []int {
	var mergedSlice []int
	if len(a) <= 1 {
		return a
	}
	if len(a) == 2 {
		if a[0] <= a[1] {
			mergedSlice = append(mergedSlice, a[0], a[1])
		} else {
			mergedSlice = append(mergedSlice, a[1], a[0])
		}
		return mergedSlice
	}
	a1 := merge(a[0:(len(a) / 2)])
	a2 := merge(a[(len(a) / 2):(len(a))])
	i := 0
	for {
		if len(a1) == 0 {
			mergedSlice = append(mergedSlice, a2...)
			break
		}
		if len(a2) == 0 {
			mergedSlice = append(mergedSlice, a1...)
			break
		}
		if a1[i] <= a2[i] {
			mergedSlice = append(mergedSlice, a1[i])
			a1 = a1[1:]
		} else {
			mergedSlice = append(mergedSlice, a2[i])
			a2 = a2[1:]
		}
	}
	return mergedSlice
}

func mergeSort(a []int) []int {
	defer libs.TimeTrack(time.Now(), "mergeSort")
	mergedSlice := merge(a)
	return mergedSlice
}

func main() {
	fmt.Println(mergeSort(libs.GenerateSlice(50)))
}
