// https://leetcode.com/problems/coin-change/
package dp

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := range dp {
		dp[i] = (amount + 1)
	}

	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if i-c >= 0 {
				dp[i] = Min(dp[i], 1+dp[i-c])
			}
		}
	}

	if dp[amount] == (amount + 1) {
		return -1
	} else {
		return dp[amount]
	}
}
