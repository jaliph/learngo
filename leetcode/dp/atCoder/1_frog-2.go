// https://atcoder.jp/contests/dp/tasks/dp_a

package dp

func Frog_2(n, k int, stones []int) int {
	Abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	const INF int = 1e9
	dp := make([]int, n)
	for i := range dp {
		dp[i] = INF
	}
	dp[0] = 0
	// fmt.Println(dp)

	for i := 1; i < n; i++ {
		for j := i - 1; j >= i-k; j-- {
			if j >= 0 {
				dp[i] = Min(dp[i], Abs(stones[i]-stones[j])+dp[j])
			}
		}
		// fmt.Println(dp)
	}
	return dp[n-1]
}
