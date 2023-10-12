package dp

// O k^n
func countWays(n, k int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}
	ans := 0
	for i := 1; i <= k; i++ {
		ans += countWays(n-i, k)
	}
	return ans
}

// top Down n*k memoise
func countWaysTD(n, k int) int {
	memoise := make([]int, 40)

	var countWays func(n, k int, dp *[]int) int
	countWays = func(n, k int, memoise *[]int) int {
		if n == 0 {
			return 1
		}
		if n < 0 {
			return 0
		}
		if (*memoise)[n] != 0 {
			return (*memoise)[n]
		}

		ans := 0
		for i := 1; i <= k; i++ {
			ans += countWays(n-i, k, memoise)
		}
		(*memoise)[n] = ans
		return ans
	}

	return countWays(n, k, &memoise)
}

func countWaysBU(n, k int) int {
	memoise := make([]int, 40)

	memoise[0] = 1

	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			if i-j >= 0 {
				memoise[i] += memoise[i-j]
			}
		}
	}
	return memoise[n]
}

// On+k
// dp[i] = dp[i - 1] - (dp[i - (k + 1)] - dp[i - 1])
// dp[i] = 2*dp[i - 1] - dp[i - k  - 1]
func countWaysOpt(n, k int) int {
	memoise := make([]int, 40)

	memoise[0] = 1
	memoise[1] = 1

	// do this for all k as dp[i - k - 1] will be less than 0
	for i := 2; i <= k; i++ {
		memoise[i] = 2 * memoise[i-1]
	}

	for i := k + 1; i <= n; i++ {
		memoise[i] = (2 * memoise[i-1]) - memoise[i-k-1]
	}

	return memoise[n]
}
