package trees

// https://leetcode.com/problems/find-mode-in-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func FindMode(root *TreeNode) []int {
	result := []int{}
	maxStreak := 0
	currStreak := 0
	var currMode int = -1e12

	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)

		// find modes
		if node.Val == currMode {
			currStreak++
		} else {
			currStreak = 0
			currMode = node.Val
		}

		if currStreak == maxStreak {
			result = append(result, currMode)
		} else if maxStreak < currStreak {
			maxStreak = currStreak
			result = []int{currMode}
		}

		inorder(node.Right)
	}

	inorder(root)
	return result
}

func FindMode_Iterative(root *TreeNode) []int {
	result := []int{}
	maxStreak := 0
	currStreak := 0
	currMode := 0
	stack := []*TreeNode{}

	curr := root

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		// pop for processing
		popped := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// find modes
		if popped.Val == currMode {
			currStreak++
		} else {
			currStreak = 1
			currMode = popped.Val
		}

		if currStreak == maxStreak {
			result = append(result, currMode)
		} else if maxStreak < currStreak {
			maxStreak = currStreak
			result = []int{currMode}
		}

		// check the right node
		curr = popped.Right
	}
	return result
}
