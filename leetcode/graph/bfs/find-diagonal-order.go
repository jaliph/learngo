package graph

// https://leetcode.com/problems/diagonal-traverse-ii/
func FindDiagonalOrder(nums [][]int) []int {
	R := len(nums)
	res := []int{}

	q := [][2]int{}
	q = append(q, [2]int{0, 0})
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		res = append(res, nums[curr[0]][curr[1]])

		if curr[1] == 0 && curr[0]+1 < R {
			q = append(q, [2]int{curr[0] + 1, 0})
		}

		if curr[1]+1 < len(nums[curr[0]]) {
			q = append(q, [2]int{curr[0], curr[1] + 1})
		}
	}

	return res
}
