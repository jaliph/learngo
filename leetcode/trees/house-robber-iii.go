// https://leetcode.com/problems/house-robber-iii/
package trees

/**
 * Definition for a binary tree node.
 */
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func Rob(root *TreeNode) int {
	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var solve func(*TreeNode) (int, int)

	// first entry denotes rob this level
	// second entry denotes rob without this level
	solve = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}

		robLeft, skipLeft := solve(node.Left)
		robRight, skipRight := solve(node.Right)

		robRoot := node.Val + skipLeft + skipRight
		skipRoot := Max(robLeft, skipLeft) + Max(robRight, skipRight)
		return robRoot, skipRoot
	}

	robRoot, skipRoot := solve(root)
	return Max(robRoot, skipRoot)
}

func Rob_1(root *TreeNode) int {
	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var solve func(*TreeNode) [2]int

	// first entry denotes rob this level
	// second entry denotes rob without this level
	solve = func(node *TreeNode) [2]int {
		if node == nil {
			return [2]int{0, 0}
		}

		left := solve(node.Left)
		right := solve(node.Right)

		robThisLevel := node.Val + left[1] + right[1]
		robWithoutLevel := Max(left[0], left[1]) + Max(right[0], right[1])
		return [2]int{robThisLevel, robWithoutLevel}
	}

	res := solve(root)
	return Max(res[0], res[1])
}
