// https://leetcode.com/problems/maximum-difference-between-node-and-ancestor/

package trees

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxAncestorDiff(root *TreeNode) int {
	const MAX int = -1e9
	const MIN int = 1e9
	var ans int = MAX

	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}
	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	Abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	fx := map[*TreeNode]int{} // min
	gx := map[*TreeNode]int{} // max

	var solve func(node *TreeNode)

	solve = func(node *TreeNode) {
		if node == nil {
			return
		}

		fx[node] = node.Val
		gx[node] = node.Val

		solve(node.Left)
		solve(node.Right)

		if node.Left != nil {
			fx[node] = Min(fx[node.Left], fx[node])
			gx[node] = Max(gx[node.Left], gx[node])
		}

		if node.Right != nil {
			fx[node] = Min(fx[node.Right], fx[node])
			gx[node] = Max(gx[node.Right], gx[node])
		}

		ans = Max(ans, Abs(node.Val-fx[node]))
		ans = Max(ans, Abs(node.Val-gx[node]))
		// fmt.Println(node, fx[node], gx[node], ans)
	}
	solve(root)

	return ans
}
