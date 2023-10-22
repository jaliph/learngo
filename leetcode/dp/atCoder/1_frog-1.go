// https://atcoder.jp/contests/dp/tasks/dp_a

package dp

func Frog_1(n int, stones []int) int {
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

	dp := make([]int, n)
	dp[1] = Abs(stones[1] - stones[0])
	for i := 2; i < n; i++ {
		dp[i] = Min(Abs(stones[i]-stones[i-1])+dp[i-1], Abs(stones[i]-stones[i-2])+dp[i-2])
	}
	return dp[n-1]
}
