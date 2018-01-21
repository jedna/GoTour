package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var walkNode func(t *tree.Tree, ch chan int)
	walkNode = func(t *tree.Tree, ch chan int) {
		if t != nil {
			walkNode(t.Left, ch)
			ch <- t.Value
			walkNode(t.Right, ch)
		}
	}
	walkNode(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, hasMore1 := <- ch1
		v2, hasMore2 := <- ch2

		if hasMore1 != hasMore2 {
			return false
		}

		if v1 != v2 {
			return false
		}

		if !hasMore1 {
			break
		}
	}

	return true
}

func main() {
	t1 := tree.New(3)
	t2 := tree.New(3)

	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(Same(t1, t2))
}
