// https://leetcode.com/problems/interleaving-string/
package dp

func IsInterleave_TD(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	var solve func(i, j int) bool

	dp := map[[2]int]bool{}
	solve = func(i, j int) bool {
		if i == len(s1) && j == len(s2) {
			return true
		}

		if i > len(s1) || j > len(s2) {
			return false
		}

		if val, ok := dp[[2]int{i, j}]; ok {
			return val
		}

		if i < len(s1) && s1[i] == s3[i+j] && solve(i+1, j) {
			dp[[2]int{i, j}] = true
			return true
		}

		if j < len(s2) && s2[j] == s3[i+j] && solve(i, j+1) {
			dp[[2]int{i, j}] = true
			return true
		}
		dp[[2]int{i, j}] = false
		return false
	}

	return solve(0, 0)
}

func IsInterleave_BU(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	R := len(s1) + 1
	C := len(s2) + 1
	dp := make([][]bool, R)
	rows := make([]bool, R*C)
	for i := range dp {
		dp[i] = rows[i*C : (i+1)*C]
	}

	dp[len(s1)][len(s2)] = true

	for i := len(s1) - 1; i >= 0; i-- {
		for j := len(s2) - 1; j >= 0; j-- {
			if i < len(s1) && dp[i+1][j] && s1[i] == s3[i+j] {
				dp[i][j] = true
			}
			if j < len(s2) && dp[i][j+1] && s2[j] == s3[i+j] {
				dp[i][j] = true
			}
		}
	}

	return dp[0][0]
}
