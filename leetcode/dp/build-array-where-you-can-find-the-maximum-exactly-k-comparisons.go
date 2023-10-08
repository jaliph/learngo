package dp

import "fmt"

// https://leetcode.com/problems/build-array-where-you-can-find-the-maximum-exactly-k-comparisons

func NumOfArrays(n int, m int, k int) int {
	dp := make(map[string]int)
	var solve func(idx, lenLIS, lar int) int
	solve = func(idx, lenLIS, lar int) int {
		// base case
		if idx == n {
			if lenLIS == k {
				return 1
			} else {
				return 0
			}
		}
		key := fmt.Sprintf("%d-%d-%d", idx, lenLIS, lar)
		if v, ok := dp[key]; ok {
			return v
		}

		ans := 0
		for i := 1; i <= m; i++ {
			if i > lar {
				ans += solve(idx+1, lenLIS+1, i)
			} else {
				ans += solve(idx+1, lenLIS, lar)
			}
			ans = ans % (1e9 + 7)
		}
		dp[key] = ans
		return ans
	}

	return solve(0, 0, 0)
}
