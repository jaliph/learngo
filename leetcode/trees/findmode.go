package trees

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
