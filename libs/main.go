package libs

import (
	"fmt"
	"time"
)

var (
	TestIntSlice = []int{1, 13, 88, 5, 46, 22, 31, 0, 155, 66, 44, 61, -45, 37, -6, -22, 9}
)

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Println(name, "took", elapsed, ", result: ")
}
