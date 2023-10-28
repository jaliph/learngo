package dp

// https://leetcode.com/problems/burst-balloons/

func maxCoins(nums []int) int {
	n := len(nums)
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)

	R := n + 2
	C := n + 2
	dp := make([][]int, R)
	rows := make([]int, R*C)
	for i := range dp {
		dp[i] = rows[i*C : (i+1)*C]
	}

	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for l := 1; l <= n; l++ { // window
		for i := 1; i <= n-l+1; i++ { // left
			j := i + l - 1 // right

			for k := i; k <= j; k++ {
				cost := nums[i-1]*nums[k]*nums[j+1] + dp[i][k-1] + dp[k+1][j]
				dp[i][j] = Max(dp[i][j], cost)
			}

		}
	}

	return dp[1][n]
}

// func Driver() {
// 	mats := []int{3, 1, 5, 8}
// 	fmt.Println(maxCoins(mats))
// }
