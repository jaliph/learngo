package main

import (
	"fmt"
)

type BinaryTreeNode struct {
	right *BinaryTreeNode
	left  *BinaryTreeNode
	data  int
}

type BinaryTree struct {
	root *BinaryTreeNode
}

func (b *BinaryTree) insert(data int) *BinaryTree {
	if b.root == nil {
		b.root = &BinaryTreeNode{
			right: nil,
			left:  nil,
			data:  data,
		}
	} else {
		b.root.insert(data)
	}
	return b
}

func (n *BinaryTreeNode) insert(data int) {
	if n == nil {
		return
	} else if data < n.data {
		if n.left == nil {
			n.left = &BinaryTreeNode{
				right: nil,
				left:  nil,
				data:  data,
			}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryTreeNode{
				right: nil,
				left:  nil,
				data:  data,
			}
		} else {
			n.right.insert(data)
		}
	}
}

func print(node *BinaryTreeNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("%c:%v\n", ch, node.data)
	print(node.left, ns+2, 'L')
	print(node.right, ns+2, 'R')
}

func main() {
	tree := &BinaryTree{}
	tree.insert(100).
		insert(-20).
		insert(-50).
		insert(-15).
		insert(-60).
		insert(50).
		insert(60).
		insert(55).
		insert(85).
		insert(15).
		insert(5).
		insert(-10)
	print(tree.root, 0, 'M')
}
