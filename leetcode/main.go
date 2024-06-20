package main

import "fmt"

type TNode struct {
	Left  *TNode
	Right *TNode
	val   int
}

type TNodeElement struct {
	T     *TNode
	depth int
}

func FindLowest(root *TNode) int {

	queue := []*TNodeElement{}

	queue = append(queue, &TNodeElement{T: root, depth: 0})
	for len(queue) > 0 {

		curr := queue[0]
		queue = queue[1:]

		if curr.T.Left == nil && curr.T.Right == nil {
			return curr.depth
		}

		if curr.T.Left != nil {
			queue = append(queue, &TNodeElement{T: curr.T.Left, depth: curr.depth + 1})
		}

		if curr.T.Right != nil {
			queue = append(queue, &TNodeElement{T: curr.T.Right, depth: curr.depth + 1})
		}
	}

	return -1
}

func main() {
	root := &TNode{val: 1}
	root.Left = &TNode{val: 2}
	root.Right = &TNode{val: 3}
	root.Left.Left = &TNode{val: 4}
	// root.Left.Right = &TNode{val: 5}

	fmt.Println(FindLowest(root))
}

// do it the other way
