package trees

// https://leetcode.com/problems/leaf-similar-trees

/* Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func LeafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	l1 := getLeafSequence(root1)
	l2 := getLeafSequence(root1)

	if len(l1) != len(l2) {
		return false
	}

	for i, v := range l1 {
		if l2[i] != v {
			return false
		}
	}

	return true
}

func getLeafSequence(node *TreeNode) []int {
	leafSequence := []int{}

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)
		if node.Left == nil && node.Right == nil {
			leafSequence = append(leafSequence, node.Val)
		}
		dfs(node.Right)
	}
	// call the function
	dfs(node)
	return leafSequence
}
