package graph

// https://leetcode.com/problems/path-sum-ii/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}

	var dfs func(*TreeNode, int, []int)
	dfs = func(node *TreeNode, sum int, path []int) {
		if node == nil {
			return
		}

		if sum < 0 {
			return
		}

		if node.Left == nil && node.Right == nil {
			if targetSum-node.Val == 0 {
				res = append(res, path)
			}
		}

		dfs(node.Left, targetSum-node.Val, append(path, node.Val))
		dfs(node.Right, targetSum-node.Val, append(path, node.Val))

	}

	dfs(root, targetSum, []int{})
	return res
}
