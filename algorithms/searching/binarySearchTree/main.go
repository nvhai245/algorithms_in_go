package main

import (
	"fmt"
	"github.com/nvhai245/algorithms_in_go/libs"
	"time"
)

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(k int) {
	if n.Key < k {
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key >= k {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

func (n *Node) Search(k int) (bool, *Node) {
	if n == nil {
		return false, nil
	}
	if n.Key < k {
		return n.Right.Search(k)
	}
	if n.Key > k {
		return n.Left.Search(k)
	}
	return true, n
}

func makeBinarySearchTree(a []int) *Node {
	defer libs.TimeTrack(time.Now(), "make binarySearchTree")
	n := &Node{}
	for _, v := range a {
		n.Insert(v)
	}
	return n
}

func main() {
	n := makeBinarySearchTree(libs.GenerateSlice(100))
	fmt.Printf("%#v", n)
	fmt.Println(n.Search(10))
}
