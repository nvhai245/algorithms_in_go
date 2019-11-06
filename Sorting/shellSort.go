package main

import "fmt"

func shellSort(a []int) []int {
	h := 1
	for h < len(a) / 3 {
		h = h * 3 + 1
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
	a := []int {1, 13, 88, 5, 46, 22, 31, 0, 155, 66, 44, 61, -45, 37, -6, -22}
	fmt.Println(shellSort(a))
}