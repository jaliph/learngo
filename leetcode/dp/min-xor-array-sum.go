// https://leetcode.com/problems/minimum-xor-sum-of-two-arrays
package dp

import "math"

func minimumXORSum(nums1 []int, nums2 []int) int {
	r, c := len(nums1), len(nums2)

	xor := make([][]int, r)
	rows := make([]int, r*c)
	for i := range xor {
		xor[i] = rows[i*c : (i+1)*c]
	}

	for i, v1 := range nums1 {
		for j, v2 := range nums2 {
			xor[i][j] = v1 ^ v2
		}
	}

	Min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	dp := map[[2]int]int{}
	var solve func(int, int) int
	solve = func(i, mask int) int {
		if i == r {
			return 0
		}
		if v, ok := dp[[2]int{i, mask}]; ok {
			return v
		}

		ans := math.MaxInt
		for j := 0; j < c; j++ {
			if (mask>>j)&1 > 0 {
				ans = Min(ans, xor[i][j]+solve(i+1, mask^(1<<j)))
			}
		}
		dp[[2]int{i, mask}] = ans
		return ans
	}

	return solve(0, (1<<c)-1)
}
