package dp

import (
	"math"
)

func MaxDotProductDP(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := [][]int{}

	// setting up memoise 2d array
	r := m + 1
	c := n + 1
	rows := make([]int, r*c)
	for i := 0; i < len(rows); i++ {
		rows[i] = math.MinInt
	}
	for i := 0; i < r; i++ {
		dp = append(dp, rows[i*c:(i+1)*c])
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			curr := nums1[i-1] * nums2[j-1]
			ans1 := curr + MaxIntV(dp[i-1][j-1], 0)
			ans2 := dp[i-1][j]
			ans3 := dp[i][j-1]
			dp[i][j] = MaxIntV(ans1, ans2, ans3)
		}
	}
	return dp[m][n]
}

func MaxDotProductMemoise(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := [][]int{}

	// setting up memoise 2d array
	r := m + 1
	c := n + 1
	rows := make([]int, r*c)
	for i := 0; i < len(rows); i++ {
		rows[i] = math.MinInt
	}
	for i := 0; i < r; i++ {
		dp = append(dp, rows[i*c:(i+1)*c])
	}

	var solve func(i, j int) int
	solve = func(i, j int) int {
		// base case
		if i == m || j == n {
			return math.MinInt
		}

		if dp[i][j] != math.MinInt {
			return dp[i][j]
		}

		curr := nums1[i] * nums2[j]
		ans1 := curr + MaxIntV(solve(i+1, j+1), 0)
		ans2 := solve(i+1, j)
		ans3 := solve(i, j+1)

		dp[i][j] = MaxIntV(ans1, ans2, ans3)
		return dp[i][j]
	}

	return solve(0, 0)
}

func MaxDotProductBrute(nums1 []int, nums2 []int) int {
	var solve func(i, j int) int

	m, n := len(nums1), len(nums2)
	solve = func(i, j int) int {
		// base case
		if i == m || j == n {
			return math.MinInt
		}

		curr := nums1[i] * nums2[j]
		ans1 := curr + MaxIntV(solve(i+1, j+1), 0)
		ans2 := solve(i+1, j)
		ans3 := solve(i, j+1)

		return MaxIntV(ans1, ans2, ans3)
	}

	return solve(0, 0)
}

func MaxIntV(args ...int) int {
	max := args[0]

	for _, v := range args {
		if max < v {
			max = v
		}
	}
	return max
}
