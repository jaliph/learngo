package dp

// https://leetcode.com/problems/k-inverse-pairs-array/

func KInversePairs(n int, k int) int {
	MOD := 1000000007
	dp := make([]int, k+1)
	dp[0] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= k; j++ {
			dp[j] = (dp[j] + dp[j-1]) % MOD
		}
		for j := k; j >= i; j-- {
			dp[j] = (dp[j] - dp[j-i] + MOD) % MOD
		}
	}
	return dp[k]
}

func KInversePairsTLE(n int, k int) int {
	const MOD int = 1e9 + 2
	dp := map[[2]int]int{}
	var recur func(int, int) int
	recur = func(n, k int) int {
		if k <= 0 {
			if k == 0 {
				return 1
			}
			return 0
		}

		ans := 0
		for i := 0; i < n; i++ {
			ans += recur(n-1, k-i)
			ans %= MOD
		}

		dp[[2]int{n, k}] = ans
		return ans
	}

	return recur(n, k)
}
