package dp

// https://leetcode.com/problems/burst-balloons/

func MaxCoins(nums []int) int {
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

/*
1 a , b, c, d,  1

	l         r
	nums[l - 1] * nums[i] * nums[r + 1]
	+ dp[l][i - 1]
	+ dp[i + 1][r]
*/
func MaxCoinsTD(nums []int) int {
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)

	dp := map[[2]int]int{}

	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var recur func(int, int) int
	recur = func(l, r int) int {
		if l > r {
			// doesnt add to our value
			return 0
		}
		if val, ok := dp[[2]int{l, r}]; ok {
			return val
		}

		total := 0
		for i := l; i <= r; i++ {
			res := nums[l-1] * nums[i] * nums[r+1]
			res += recur(l, i-1)
			res += recur(i+1, r)
			total = Max(total, res)
		}
		dp[[2]int{l, r}] = total
		return total
	}

	return recur(1, len(nums)-2)
}
