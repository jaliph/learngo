package dp

import "fmt"

func MCM(mats []int) int {

	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	R := len(mats)
	C := len(mats)
	dp := make([][]int, R)
	rows := make([]int, R*C)
	for i := range rows {
		rows[i] = -1
	}
	for i := range dp {
		dp[i] = rows[i*C : (i+1)*C]
	}

	Print2D := func(grid [][]int) {
		for r := range grid {
			for _, colValue := range grid[r] {
				fmt.Print(colValue, " ")
			}
			fmt.Println()
		}
		fmt.Println("----")
	}

	var solve func(int, int) int
	solve = func(i, j int) int {
		if i >= j {
			dp[i][j] = 0
			Print2D(dp)
			return 0
		}

		if dp[i][j] != -1 {
			return dp[i][j]
		}

		var temp int = 1e9
		for k := i; k <= j-1; k++ {
			left := solve(i, k)
			right := solve(k+1, j)
			cost := mats[i-1] * mats[k] * mats[j]
			temp = Min(temp, cost+left+right)
		}
		dp[i][j] = temp
		Print2D(dp)
		return temp
	}

	res := solve(1, len(mats)-1)
	return res
}

func MCM_DP(mats []int) int {
	n := len(mats)

	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	R := n
	C := n
	dp := make([][]int, R)
	rows := make([]int, R*C)
	for i := range dp {
		dp[i] = rows[i*C : (i+1)*C]
	}

	for l := 2; l < n; l++ { // window
		for i := 1; i < n-l+1; i++ { // left
			j := i + l - 1 // right
			dp[i][j] = 1e9
			for k := i; k < j; k++ {
				cost := mats[i-1]*mats[k]*mats[j] + dp[i][k] + dp[k+1][j]
				dp[i][j] = Min(dp[i][j], cost)
			}
		}
	}
	return dp[1][n-1]
}

// func Driver() {
// 	mats := []int{40, 20, 30, 10, 30}
// 	fmt.Println(MCM_DP(mats))
// }
