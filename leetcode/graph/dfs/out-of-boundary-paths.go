package graph

// https://leetcode.com/problems/out-of-boundary-paths/description/
func FindPaths(m int, n int, maxMove int, startRow int, startColumn int) int {

	const MOD int = 1e9 + 7
	R := m
	C := n

	paths := [][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	dp := map[[2]int]int{}
	var dfs func(i, j, leftMoves int) int
	dfs = func(i, j, leftMoves int) int {
		// fmt.Println(i, j, leftMoves)
		// base
		if i < 0 || j < 0 || i >= R || j >= C {
			if leftMoves >= 0 {
				// fmt.Println("Possible ans ", i, j, leftMoves)
				return 1
			}
			// fmt.Println("Possible not ans ", i, j, leftMoves)
			return 0
		}

		if leftMoves <= 0 {
			return 0
		}

		if v, ok := dp[[2]int{i, j}]; ok {
			return v
		}

		// recur
		count := 0
		for _, v := range paths {
			count += dfs(i+v[0], j+v[1], leftMoves-1)
		}

		dp[[2]int{i, j}] = count
		return count
	}
	val := dfs(startRow, startColumn, maxMove)
	// fmt.Println(dp)
	return val
}

func findPathsDP(m int, n int, maxMove int, startRow int, startColumn int) int {

	memoise := make([][][]int, maxMove+1)
	for i := 0; i < maxMove+1; i++ {
		R := m
		C := n
		dp := make([][]int, R)
		rows := make([]int, R*C)
		for i := range dp {
			dp[i] = rows[i*C : (i+1)*C]
		}
		memoise[i] = dp
	}

	dirs := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	const MOD int = 1e9 + 7
	for leftmoves := 1; leftmoves <= maxMove; leftmoves++ {

		for r := 0; r < m; r++ {

			for c := 0; c < n; c++ {
				ans := 0
				for _, v := range dirs {
					n_x := r + v[0]
					n_y := c + v[1]
					if n_x < 0 || n_y < 0 || n_x >= m || n_y >= n {
						ans++
					} else {
						ans += memoise[leftmoves-1][n_x][n_y]
					}
				}
				ans %= MOD
				memoise[leftmoves][r][c] = ans
			}
		}
	}

	return memoise[maxMove][startRow][startColumn]
}

func findPathsDP_Opt(m int, n int, maxMove int, startRow int, startColumn int) int {

	memoise := make([][][]int, 2)
	for i := 0; i < 2; i++ {
		R := m
		C := n
		dp := make([][]int, R)
		rows := make([]int, R*C)
		for i := range dp {
			dp[i] = rows[i*C : (i+1)*C]
		}
		memoise[i] = dp
	}

	dirs := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	const MOD int = 1e9 + 7
	for leftmoves := 1; leftmoves <= maxMove; leftmoves++ {

		for r := 0; r < m; r++ {
			for c := 0; c < n; c++ {
				ans := 0
				for _, v := range dirs {
					n_x := r + v[0]
					n_y := c + v[1]
					if n_x < 0 || n_y < 0 || n_x >= m || n_y >= n {
						ans++
					} else {
						ans += memoise[(leftmoves-1)&1][n_x][n_y]
					}
				}
				ans %= MOD
				memoise[leftmoves&1][r][c] = ans
			}
		}
	}

	return memoise[maxMove&1][startRow][startColumn]
}

// func Driver() {
// 	m := 2
// 	n := 2
// 	maxMove := 2
// 	startRow := 0
// 	startColumn := 0

// 	fmt.Println(findPaths(m, n, maxMove, startRow, startColumn))

// 	m = 1
// 	n = 3
// 	maxMove = 3
// 	startRow = 0
// 	startColumn = 1
// 	fmt.Println(findPaths(m, n, maxMove, startRow, startColumn))
// }
