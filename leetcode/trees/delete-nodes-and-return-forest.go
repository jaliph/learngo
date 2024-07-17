package trees

// https://leetcode.com/problems/delete-nodes-and-return-forest

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func DelNodes(root *TreeNode, to_delete []int) []*TreeNode {

	toDeleteSet := map[int]struct{}{}
	res := map[*TreeNode]struct{}{
		root: {},
	}

	for _, v := range to_delete {
		toDeleteSet[v] = struct{}{}
	}

	var recur func(*TreeNode, *TreeNode, bool)
	recur = func(node *TreeNode, par *TreeNode, isLeft bool) {
		if node == nil {
			return
		}
		recur(node.Left, node, true)
		recur(node.Right, node, false)

		if _, ok := toDeleteSet[node.Val]; ok {
			delete(res, node)
			if par != nil {
				if isLeft {
					par.Left = nil
				} else {
					par.Right = nil
				}
			}
			if node.Left != nil {
				res[node.Left] = struct{}{}
			}
			if node.Right != nil {
				res[node.Right] = struct{}{}
			}
		}
	}

	recur(root, nil, false)

	result := []*TreeNode{}
	for k := range res {
		result = append(result, k)
	}
	return result
}
