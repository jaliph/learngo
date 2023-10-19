// https://leetcode.com/problems/coin-change/
package dp

func Change(amount int, coins []int) int {
	prev := make([]int, amount+1)
	prev[0] = 1

	for i := len(coins) - 1; i >= 0; i-- {
		dp := make([]int, amount+1)
		dp[0] = 1
		for j := 1; j <= amount; j++ {
			if i+1 < len(coins) {
				dp[j] = prev[j]
			}
			if j-coins[i] >= 0 {
				dp[j] += dp[j-coins[i]]
			}
		}

		prev = dp
	}
	return prev[amount]
}

func Change_2D_DP(amount int, coins []int) int {
	R, C := len(coins), amount+1
	dp := make([][]int, R)

	rows := make([]int, R*C)

	for i := range dp {
		dp[i] = rows[i*C : (i+1)*C]
	}

	for i := 0; i < R; i++ {
		dp[i][0] = 1
	}

	for i := R - 1; i >= 0; i-- {
		for j := 1; j <= amount; j++ {
			if i+1 < R {
				dp[i][j] = dp[i+1][j]
			}
			if j-coins[i] >= 0 {
				dp[i][j] += dp[i][j-coins[i]]
			}
		}
	}

	return dp[0][amount]
}
