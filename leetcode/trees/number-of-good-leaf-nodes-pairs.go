package trees

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// https://leetcode.com/problems/number-of-good-leaf-nodes-pairs/?envType=daily-question&envId=2024-07-18
func countPairs(root *TreeNode, distance int) int {
	cnt := 0

	var dfs func(*TreeNode) []int
	dfs = func(tn *TreeNode) []int {
		if tn == nil {
			return []int{}
		}
		if tn.Left == nil && tn.Right == nil {
			return []int{1}
		}

		left := dfs(tn.Left)
		right := dfs(tn.Right)

		for _, l := range left {
			for _, r := range right {
				if l+r <= distance {
					cnt++
				}
			}
		}

		final := append(left, right...)
		for i := range final {
			final[i]++
		}
		return final
	}
	dfs(root)
	return cnt
}
