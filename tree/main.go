package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

/*
type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}
*/

func WalkInOrder(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	var nodes []*tree.Tree
	node := t
	for node != nil || len(nodes) > 0 {
		for node != nil {
			nodes = append(nodes, node)
			node = node.Left
		}

		last_idx := len(nodes) - 1
		node = nodes[last_idx]
		nodes = nodes[:last_idx]

		ch <- node.Value
		node = node.Right
	}
}

func WalkBFS(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	var nodes []*tree.Tree
	var node *tree.Tree

	nodes = append(nodes, t)

	for len(nodes) > 0 {
		last_idx := len(nodes) - 1
		node = nodes[last_idx]
		ch <- node.Value
		nodes = nodes[:last_idx]

		if node.Left != nil {
			nodes = append(nodes, node.Left)
		}

		if node.Right != nil {
			nodes = append(nodes, node.Right)

		}
	}

}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func WalkRecursive(node *tree.Tree, ch chan int) {

	if node == nil {
		return
	}

	if node.Left != nil {
		WalkRecursive(node.Left, ch)
	}

	ch <- node.Value

	if node.Right != nil {
		WalkRecursive(node.Right, ch)
	}

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree, walker func(node *tree.Tree, ch chan int)) bool {
	const n int = 10
	ch1 := make(chan int)
	ch2 := make(chan int)
	go walker(t1, ch1)
	go walker(t2, ch2)

	for i := 0; i < n; i++ {
		val1 := <-ch1
		val2 := <-ch2
		if val1 != val2 {
			return false
		}
	}
	return true
}

func main() {

	fmt.Printf("WalkRecursive: Value same? true == %v\n", Same(tree.New(1), tree.New(1), WalkRecursive))
	fmt.Printf("WalkRecursive: Value same? false == %v\n", Same(tree.New(1), tree.New(2), WalkRecursive))

	fmt.Printf("WalkInOrder: Value same? true == %v\n", Same(tree.New(1), tree.New(1), WalkInOrder))
	fmt.Printf("WalkInOrder: Value same? false == %v\n", Same(tree.New(1), tree.New(2), WalkInOrder))

	fmt.Printf("WalkBFS: Value same? true == %v\n", Same(tree.New(1), tree.New(1), WalkBFS))
	fmt.Printf("WalkBFS: Value same? false == %v\n", Same(tree.New(1), tree.New(2), WalkBFS))
}
