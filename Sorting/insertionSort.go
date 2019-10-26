package main

import "fmt"

func insertionSort(a []int) []int {
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
	a := []int {1, 13, 88, 5, 46, 22, 31, 0, 155, 66, 44, 61, -45, 37, -6, -22}
	fmt.Println(insertionSort(a))
}