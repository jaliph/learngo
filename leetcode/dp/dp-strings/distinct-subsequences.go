package dp

// https://leetcode.com/problems/distinct-subsequences/description/

func NumDistinct_TD(s string, t string) int {

	dp := map[[2]int]int{}

	var count func(int, int) int
	count = func(i, j int) int {
		if j == len(t) {
			return 1
		}
		if i == len(s) {
			return 0
		}

		if c, ok := dp[[2]int{i, j}]; ok {
			return c
		}

		res := 0
		if s[i] == t[j] {
			res += count(i+1, j+1)
			res += count(i+1, j)
		} else {
			res += count(i+1, j)
		}
		dp[[2]int{i, j}] = res
		return res
	}

	return count(0, 0)
}

func NumDistinct_BU(s string, t string) int {

	R := len(s) + 1
	C := len(t) + 1
	dp := make([][]int, R)
	rows := make([]int, R*C)
	for i := range dp {
		dp[i] = rows[i*C : (i+1)*C]
	}

	for i := 0; i <= len(s); i++ {
		dp[i][len(t)] = 1
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := len(t) - 1; j >= 0; j-- {
			if s[i] == t[j] {
				dp[i][j] += dp[i+1][j+1] + dp[i+1][j]
			} else {
				dp[i][j] += dp[i+1][j]
			}
		}
	}
	return dp[0][0]
}
