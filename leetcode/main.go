package main

import (
	"fmt"
)

func main() {
	// m := map[int]bool{}
	// fmt.Println(m[2])

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	val int
	x   int
	y   int
}

func verticalTraversal(root *TreeNode) [][]int {

	// const INF int = 1001
	// min := INF
	// max := -INF

	m := map[Node]bool{}

	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, i, j int) {
		if root == nil {
			return
		}

		m[Node{node.Val, i, j}] = true
		dfs(node.Left, i+1, j-1)
		dfs(node.Right, i+1, j+1)
	}

	dfs(root, 0, 0)

	fmt.Println(m)
	return [][]int{}
}

// do it the other way
