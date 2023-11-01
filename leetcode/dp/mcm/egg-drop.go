package dp

// TLE
// https://leetcode.com/problems/super-egg-drop/
func SuperEggDrop(k int, n int) int {

	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	R := k + 1
	C := n + 1
	dp := make([][]int, R)
	rows := make([]int, R*C)
	for i := range rows {
		rows[i] = 1e9
	}
	for i := range dp {
		dp[i] = rows[i*C : (i+1)*C]
	}
	var solve func(int, int) int
	solve = func(e, f int) int {
		if f == 1 || f == 0 {
			return f
		}

		if e == 1 {
			return f
		}

		if dp[e][f] != 1e9 {
			return dp[e][f]
		}

		var ans int = 1e9
		for i := 1; i <= f; i++ {
			temp := 1 + Max(solve(e-1, i-1), solve(e, f-i))
			ans = Min(ans, temp)
		}
		dp[e][f] = ans
		return ans
	}

	return solve(k, n)
}

/// the above boils down to the formlae
//  dp[m][e] = 1 + dp[m - 1][e - 1] + dp[m - 1][e]
// where dp[m][e] represents m moves and e eggs

func SuperEggDrop_V2(e int, f int) int {
	R := f + 1
	C := e + 1
	dp := make([][]int, R)
	rows := make([]int, R*C)
	for i := range dp {
		dp[i] = rows[i*C : (i+1)*C]
	}

	m := 0 // moves
	for dp[m][e] < f {
		m++
		for x := 1; x <= e; x++ {
			dp[m][x] = 1 + dp[m-1][x-1] + dp[m-1][x]
		}
	}
	return m
}

// Space optimised
func SuperEggDrop_V3(e int, f int) int {

	dp := make([]int, e+1)
	m := 0 // moves
	for dp[e] < f {
		m++
		for x := e; x >= e; x-- {
			dp[x] = 1 + dp[x-1] + dp[x]
		}
	}
	return m
}

// func Driver() {
// 	fmt.Println(superEggDrop_V2(1, 10))
// }
