// https://leetcode.com/problems/stone-game/
package dp

func StoneGame(piles []int) bool {
	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	n := len(piles)
	dp := make([][]int, n)
	rows := make([]int, n*n)
	for i := range rows {
		rows[i] = -1
	}

	for i := range dp {
		dp[i] = rows[i*n : (i+1)*n]
	}

	var solve func(int, int) int
	solve = func(l, r int) int {
		if l == r {
			return piles[l]
		}

		if dp[l][r] != -1 {
			return dp[l][r]
		}

		res := Max(piles[l]-solve(l+1, r), piles[r]-solve(l, r-1))

		dp[l][r] = res
		return res
	}

	return solve(0, n-1) > 0
}
