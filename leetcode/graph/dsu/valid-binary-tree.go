package graph

import "fmt"

// https://leetcode.com/problems/validate-binary-tree-nodes/

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {

	d := NewDSU(n)

	for i := 0; i < n; i++ {
		if leftChild[i] != -1 {
			if d.Find(i) == d.Find(leftChild[i]) {
				return false
			} else {
				d.Union(i, leftChild[i])
			}
		}

		if rightChild[i] != -1 {
			if d.Find(i) == d.Find(rightChild[i]) {
				return false
			} else {
				d.Union(i, rightChild[i])
			}
		}
	}

	fmt.Println(d.parents)

	return d.IsConnected()
}

// func Driver() {
// 	n := 3
// 	leftChild := []int{1, -1, -1}
// 	rightChild := []int{-1, -1, 1}

// 	fmt.Println(validateBinaryTreeNodes(n, leftChild, rightChild))
// }
