// https://leetcode.com/problems/amount-of-time-for-binary-tree-to-be-infected/

// https://leetcode.com/problems/amount-of-time-for-binary-tree-to-be-infected/editorial

package trees

func AmountOfTime(root *TreeNode, start int) int {
	ans := 0
	trav(root, start, &ans)
	return ans
}

func trav(node *TreeNode, start int, ans *int) int {
	if node == nil {
		return 0
	}

	lDist := trav(node.Left, start, ans)
	rDist := trav(node.Right, start, ans)

	if node.Val == start {
		*ans = Max(*ans, Max(lDist, rDist))
		return -1
	}

	// ans not
	if lDist >= 0 && rDist >= 0 {
		return Max(lDist, rDist) + 1
	}

	*ans = Max(*ans, Abs(lDist)+Abs(rDist))

	return Min(lDist, rDist) - 1
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
