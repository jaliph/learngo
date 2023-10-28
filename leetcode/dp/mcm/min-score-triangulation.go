// https://leetcode.com/problems/minimum-score-triangulation-of-polygon/solutions/286753/c-with-picture/

package dp

func MinScoreTriangulation(values []int) int {
	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}
	const INF int = 1e15
	n := len(values)
	R := n
	C := n
	dp := make([][]int, R)
	rows := make([]int, R*C)
	for i := range dp {
		dp[i] = rows[i*C : (i+1)*C]
	}

	for i := n - 2; i >= 0; i++ {
		for j := i + 2; j < n; j++ {
			dp[i][j] = INF
			for k := i + 1; k < j; k++ {
				dp[i][j] = Min(dp[i][j], dp[i][k]+(values[i]*values[k]*values[j])+dp[k][j])
				// make two points, check for all the third point.
				// dp[i][k] & dp[k][j] represent the polgons on the left and right
			}
		}
	}

	return dp[0][n-1]
}
