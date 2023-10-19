package dp

import (
	"fmt"
	"math"
)

// Binary Search
func CalculateMinimumHP(dungeon [][]int) int {
	R, C := len(dungeon), len(dungeon[0])

	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	isAlive := func(init int, dungeon [][]int) bool {

		dp := make([][]int, R)
		rows := make([]int, R*C)
		for i := range dp {
			dp[i] = rows[i*C : (i+1)*C]
		}

		dp[0][0] = dungeon[0][0] + init

		for i := 0; i < R; i++ {
			for j := 0; j < C; j++ {
				if i > 0 && dp[i-1][j] > 0 {
					dp[i][j] = Max(dp[i][j], dungeon[i][j]+dp[i-1][j])
				}

				if j > 0 && dp[i][j-1] > 0 {
					dp[i][j] = Max(dp[i][j], dungeon[i][j]+dp[i][j-1])
				}
			}
		}
		return dp[R-1][C-1] > 0
	}

	l, r := 1, (R+C)+1000+1
	var mid int
	ans := r
	for l < r {
		mid = l + (r-l)/2
		if isAlive(mid, dungeon) {
			ans = mid
			r = mid
		} else {
			l = mid + 1
		}
	}

	return ans
}

// DP
func CalculateMinimumHP_DP(dungeon [][]int) int {
	r, c := len(dungeon), len(dungeon[0])
	m, n := r+1, c+1
	fmt.Println(m, n)
	grid := make([][]int, m)
	rows := make([]int, m*n)

	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := range rows {
		rows[i] = math.MaxInt
	}

	for i := range grid {
		grid[i] = rows[i*n : (i+1)*n]
	}

	grid[r][c-1] = 1
	grid[r-1][c] = 1

	for i := r - 1; i >= 0; i-- {
		for j := c - 1; j >= 0; j-- {
			need := Min(grid[i+1][j], grid[i][j+1]) - dungeon[i][j]
			if need <= 0 {
				grid[i][j] = 1
			} else {
				grid[i][j] = need
			}
		}
	}
	return grid[0][0]
}
