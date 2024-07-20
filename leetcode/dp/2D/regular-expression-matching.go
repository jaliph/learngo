package dp

func IsMatch(s string, p string) bool {
	R := len(s)
	C := len(p)

	// memoise
	rows := make([]int, (R+1)*(C+1))
	for i := range rows {
		rows[i] = -1
	}
	dp := make([][]int, (R + 1))
	for i := range dp {
		dp[i] = rows[i*C : (i+1)*C]
	}

	// states i for s, j for p
	var recur func(int, int) int
	recur = func(i, j int) int {

		if i == R && j == C {
			return 1
		}
		if j == C {
			return 0
		}

		if dp[i][j] != -1 {
			return dp[i][j]
		}

		match := i < R && (s[i] == p[j] || p[j] == '.')
		if j+1 < C && p[j+1] == '*' {
			res := 0
			// i can choose to not pick any char before j*
			res += recur(i, j+2)
			// i can choose only 1 char
			if match {
				res += recur(i+1, j)
			}
			if res > 0 {
				dp[i][j] = 1
				return 1
			}
			dp[i][j] = 0
			return 0
		}

		if match {
			res := recur(i+1, j+1)
			dp[i][j] = res
			return res
		}
		dp[i][j] = 0
		return 0
	}
	return recur(0, 0) == 1
}
