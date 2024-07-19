package dp

// https://leetcode.com/problems/coin-change
func CoinChange(coins []int, amount int) int {

	Min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for _, c := range coins {
		for i := 1; i <= amount; i++ {
			if c <= i {
				dp[i] = Min(dp[i], 1+dp[i-c])
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
