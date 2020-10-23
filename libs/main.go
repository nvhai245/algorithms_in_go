package libs

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	TestIntSlice = []int{1, 13, 88, 5, 46, 22, 31, 0, 155, 66, 44, 61, -45, 37, -6, -22, 9}
)

// Generates a slice of size, size filled with random numbers
func GenerateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	fmt.Println("[UNSORTED]:")
	fmt.Println()
	fmt.Println(slice)
	fmt.Println()
	return slice
}

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
	return
}

// Generates a slice of size, size filled with random numbers
func GenerateSortedSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	sort(slice)
	fmt.Println("[SORTED]:")
	fmt.Println()
	fmt.Println(slice)
	fmt.Println()
	return slice
}

// TimeTrack show the total execution time of a function
func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Println(name, "took", elapsed, ", result: ")
	fmt.Println()
}
