package trees

import "sort"

type PNode struct {
	val int
	x   int
}

func VerticalTraversal(root *TreeNode) [][]int {

	const INF int = 1001
	min := INF
	// max := -INF
	Min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	m := map[int][]PNode{}

	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, i, j int) {
		if node == nil {
			return
		}
		min = Min(min, j)
		m[j] = append(m[j], PNode{node.Val, i})
		dfs(node.Left, i+1, j-1)
		dfs(node.Right, i+1, j+1)
	}

	dfs(root, 0, 0)

	res := make([][]int, len(m))
	for k, v := range m {
		sort.Slice(v, func(i, j int) bool {
			if v[i].x == v[j].x {
				return v[i].val < v[j].val
			}
			return v[i].x < v[j].x
		})
		for _, node := range v {
			res[k+(-min)] = append(res[k+(-min)], node.val)
		}
	}
	return res
}
