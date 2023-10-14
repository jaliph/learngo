// https://leetcode.com/problems/perfect-squares/
package dp

func numSquares(n int) int {

	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	nums := []int{}
	j := 1
	for j*j <= n {
		nums = append(nums, j*j)
		j++
	}

	dp := make([]int, n+1)

	for i := range dp {
		dp[i] = n + 1
	}

	dp[0] = 0

	for i := 0; i <= n; i++ {
		for _, num := range nums {
			if i-num >= 0 {
				dp[i] = Min(dp[i], 1+dp[i-num])
			}
		}
	}

	return dp[n]
}
