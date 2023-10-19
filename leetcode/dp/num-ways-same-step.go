package dp

func numWays(steps int, arrLen int) int {
	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}
	const MOD = 1e9 + 7

	dp := make([][]int, steps+1)

	limit := Min(arrLen, steps/2+1) /// can go till here and come back to index if max allowed is steps

	rows := make([]int, (steps+1)*limit)
	for i := range dp {
		dp[i] = rows[i*limit : (i+1)*limit]
	}

	dp[0][0] = 1

	for i := 1; i <= steps; i++ {
		for j := 0; j < limit; j++ {
			dp[i][j] = dp[i-1][j] // i can chose to stay, count that
			if j > 0 {
				dp[i][j] += dp[i-1][j-1] // count from left
			}
			if j < limit-1 {
				dp[i][j] += dp[i-1][j+1] // count from left
			}
			dp[i][j] %= MOD
		}
	}
	return dp[steps][0]
}

func numWays_TD(steps int, arrLen int) int {

	const MOD = 1e9 + 7

	dp := map[[2]int]int{}
	var solve func(int, int) int
	solve = func(idx, steps int) int {
		// base
		if steps == 0 {
			if idx == 0 {
				return 1
			} else {
				return 0
			}
		}

		if steps < 0 || idx < 0 || idx >= arrLen {
			return 0
		}

		if v, ok := dp[[2]int{idx, steps}]; ok {
			return v
		}

		// recur

		ways := solve(idx, steps-1)
		ways += solve(idx+1, steps-1)
		ways += solve(idx-1, steps-1)
		ways %= MOD

		dp[[2]int{idx, steps}] = ways
		return ways
	}

	return solve(0, steps)
}
