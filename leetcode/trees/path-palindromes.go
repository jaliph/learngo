package trees

// https://leetcode.com/problems/pseudo-palindromic-paths-in-a-binary-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pseudoPalindromicPaths(root *TreeNode) int {

	var dfs func(*TreeNode, int) int
	dfs = func(node *TreeNode, pathBits int) int {
		if node == nil {
			return 0
		}

		pathBits = pathBits ^ (1 << node.Val)

		if node.Left == nil && node.Right == nil {
			if pathBits&(pathBits-1) == 0 {
				return 1
			}
			return 0
		}
		return dfs(node.Left, pathBits) + dfs(node.Right, pathBits)
	}

	return dfs(root, 0)
}
