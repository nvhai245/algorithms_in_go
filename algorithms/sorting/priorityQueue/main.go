package main

import (
	"fmt"
	"github.com/nvhai245/algorithms_in_go/libs"
	"time"
)

type MaxHeap struct {
	slice []int
}

// Insert element to a heap
func (h *MaxHeap) Insert(num int) {
	h.slice = append(h.slice, num)
	h.heapifyUp(len(h.slice) - 1)
}

// Extract the first element from a heap
func (h *MaxHeap) Extract() int {
	if len(h.slice) == 0 {
		fmt.Println("Cannot extract from an empty heap")
		return 0
	}
	extracted := h.slice[0]
	l := len(h.slice) - 1
	h.slice[0] = h.slice[l]
	h.slice = h.slice[:l]

	h.heapifyDown(0)
	return extracted
}

// Extract all elements from a heap
func (h *MaxHeap) ExtractAll() []int {
	defer libs.TimeTrack(time.Now(), "heapSort")
	var a []int
	for len(h.slice) > 0 {
		a = append(a, h.Extract())
	}
	return a
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.slice[i1], h.slice[i2] = h.slice[i2], h.slice[i1]
}

func (h *MaxHeap) heapifyUp(index int) {
	for h.slice[parentIndex(index)] < h.slice[index] {
		h.swap(parentIndex(index), index)
		index = parentIndex(index)
	}
}

func (h *MaxHeap) heapifyDown(index int) {
	last := len(h.slice) - 1
	left, right := leftIndex(index), rightIndex(index)
	child := 0
	for left <= last {
		if left == last {
			child = left
		} else if h.slice[left] > h.slice[right] {
			child = left
		} else {
			child = right
		}
		if h.slice[index] < h.slice[child] {
			h.swap(index, child)
			index = child
			left, right = leftIndex(index), rightIndex(index)
		} else {
			return
		}
	}
}

func parentIndex(i int) int {
	return (i - 1) / 2
}

func leftIndex(i int) int {
	return 2*i + 1
}

func rightIndex(i int) int {
	return 2*i + 2
}

func makeHeap(a []int) *MaxHeap {
	defer libs.TimeTrack(time.Now(), "makeMaxHeap")
	m := &MaxHeap{}
	for _, num := range a {
		m.Insert(num)
	}
	return m
}

func main() {
	a := makeHeap(libs.GenerateSlice(0))
	fmt.Println(a.slice)
	fmt.Println(a.ExtractAll())
}
