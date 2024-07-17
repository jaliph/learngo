package trees

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/step-by-step-directions-from-a-binary-tree-node-to-another/

func GetDirections(root *TreeNode, startValue int, destValue int) string {
	var dfs func(*TreeNode, *[]string, int) bool
	dfs = func(node *TreeNode, s *[]string, dst int) bool {
		if node == nil {
			return false
		}
		if node.Val == dst {
			return true
		}

		*s = append(*s, "L")
		if dfs(node.Left, s, dst) {
			return true
		}
		*s = (*s)[:len(*s)-1]

		*s = append(*s, "R")
		if dfs(node.Right, s, dst) {
			return true
		}
		*s = (*s)[:len(*s)-1]

		return false
	}

	pathTo := []string{}
	dfs(root, &pathTo, destValue)
	pathFrom := []string{}
	dfs(root, &pathFrom, startValue)

	Min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	idx := 0
	for idx < Min(len(pathTo), len(pathFrom)) && pathTo[idx] == pathFrom[idx] {
		idx++
	}
	// fmt.Println(idx)
	out := ""
	for i := 0; i < len(pathFrom)-idx; i++ {
		out += "U"
	}

	return out + strings.Join(pathTo[idx:], "")
}

func GetDirections1(root *TreeNode, startValue int, destValue int) string {

	var LCA func(root *TreeNode, x, y int) *TreeNode
	LCA = func(root *TreeNode, x, y int) *TreeNode {
		if root == nil || root.Val == x || root.Val == y {
			return root
		}

		l := LCA(root.Left, x, y)
		r := LCA(root.Right, x, y)
		if l == nil {
			return r
		}
		if r == nil {
			return l
		}
		return root
	}

	var dfs func(*TreeNode, *[]string, int, bool) bool
	dfs = func(node *TreeNode, s *[]string, dst int, isReverse bool) bool {
		if node == nil {
			return false
		}
		if node.Val == dst {
			return true
		}

		ch := "L"
		if isReverse {
			ch = "U"
		}
		*s = append(*s, ch)
		if dfs(node.Left, s, dst, isReverse) {
			return true
		}
		*s = (*s)[:len(*s)-1]

		ch = "R"
		if isReverse {
			ch = "U"
		}
		*s = append(*s, ch)
		if dfs(node.Right, s, dst, isReverse) {
			return true
		}
		*s = (*s)[:len(*s)-1]

		return false
	}

	par := LCA(root, startValue, destValue)

	pathTo := []string{}
	dfs(par, &pathTo, destValue, false)
	pathFrom := []string{}
	dfs(par, &pathFrom, destValue, true)

	fmt.Println(pathTo, pathFrom)

	return strings.Join(pathFrom, "") + strings.Join(pathTo, "")
}
