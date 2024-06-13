package graph

// https://leetcode.com/problems/path-sum-iii/

func PathSum3(root *TreeNode, targetSum int) int {

	count := 0
	countMap := map[int]int{}
	countMap[0] = 1

	sum := 0
	var dfs func(*TreeNode, *int)
	dfs = func(node *TreeNode, sum *int) {
		if node == nil {
			return
		}

		*sum += node.Val

		if val, ok := countMap[*sum-targetSum]; ok {
			count += val
		}
		countMap[*sum]++

		dfs(node.Left, sum)
		dfs(node.Right, sum)
		countMap[*sum]--
		*sum -= node.Val
	}

	dfs(root, &sum)
	return count
}
