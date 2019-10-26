package main

import "fmt"

func selectionSort(a []int) []int {
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
	a := []int {1, 13, 88, 5, 46, 22, 31, 0, 155, 66, 44, 61, -45, 37, -6, -22}
	fmt.Println(selectionSort(a))
}