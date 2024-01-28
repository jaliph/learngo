package dp

// https://leetcode.com/problems/knight-probability-in-chessboard/

func knightProbabilityBrute(n int, k int, row int, column int) float64 {

	dp := map[[3]int]float64{}

	dir := [][]int{
		{-2, -1}, {-1, -2}, {1, -2}, {2, -1}, {2, 1}, {1, 2}, {-1, 2}, {-2, 1},
	}
	var solve func(int, int, int) float64

	solve = func(i, j, left int) float64 {
		// base
		if left < 0 {
			return 1
		}
		if i < 0 || j < 0 || i > n || j > n {
			return 0
		}

		if v, ok := dp[[3]int{i, j, left}]; ok {
			return v
		}

		// recur
		ans := 0.0
		for _, v := range dir {
			ans = ans + (solve(i+v[0], j+v[1], left-1) / float64(8))
		}
		dp[[3]int{1, j, left}] = ans
		return ans
	}

	return solve(row, column, k)
}

func knightProbabilityDP(n int, k int, row int, column int) float64 {

	memoise := make([][][]float64, k+1)

	for i := range memoise {
		R := n
		C := n
		dp := make([][]float64, R)
		rows := make([]float64, R*C)
		for i := range dp {
			dp[i] = rows[i*C : (i+1)*C]
		}
		memoise[i] = dp
	}

	memoise[0][row][column] = 1

	dir := [][]int{
		{-2, -1}, {-1, -2}, {1, -2}, {2, -1}, {2, 1}, {1, 2}, {-1, 2}, {-2, 1},
	}
	for p := 1; p <= k; p++ {

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {

				// next possible jump positions
				for _, v := range dir {
					n_i := i + v[0]
					n_j := j + v[1]

					if n_i >= 0 && n_j >= 0 && n_i < n && n_j < n {
						memoise[p][i][j] += (memoise[p-1][n_i][n_j] / float64(8.0))
					}
				}
			}
		}
	}

	ans := 0.0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ans += memoise[k][i][j]
		}
	}
	return ans
}

func knightProbabilityOpt(n int, k int, row int, column int) float64 {

	memoise := make([][][]float64, 2)

	for i := range memoise {
		R := n
		C := n
		dp := make([][]float64, R)
		rows := make([]float64, R*C)
		for i := range dp {
			dp[i] = rows[i*C : (i+1)*C]
		}
		memoise[i] = dp
	}

	memoise[0][row][column] = 1

	dir := [][]int{
		{-2, -1}, {-1, -2}, {1, -2}, {2, -1}, {2, 1}, {1, 2}, {-1, 2}, {-2, 1},
	}
	for p := 1; p <= k; p++ {

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				sum := 0.0
				for _, v := range dir {
					n_i := i + v[0]
					n_j := j + v[1]

					if n_i >= 0 && n_j >= 0 && n_i < n && n_j < n {
						sum += (memoise[(p-1)&1][n_i][n_j] / float64(8.0))
					}
				}
				memoise[p&1][i][j] = sum
			}
		}
	}

	ans := 0.0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ans += memoise[k&1][i][j]
		}
	}
	return ans
}
