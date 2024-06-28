// https://leetcode.com/problems/cherry-pickup-ii/
package graph

func cherryPickup2(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	dp := make([][][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, n)
		}
	}

	for row := m - 1; row >= 0; row-- {
		for col1 := 0; col1 < n; col1++ {
			for col2 := 0; col2 < n; col2++ {
				result := 0
				// current cell
				result += grid[row][col1]
				if col1 != col2 {
					result += grid[row][col2]
				}
				// transition
				if row != m-1 {
					max := 0
					for newCol1 := col1 - 1; newCol1 <= col1+1; newCol1++ {
						for newCol2 := col2 - 1; newCol2 <= col2+1; newCol2++ {
							if newCol1 >= 0 && newCol1 < n && newCol2 >= 0 && newCol2 < n {
								if dp[row+1][newCol1][newCol2] > max {
									max = dp[row+1][newCol1][newCol2]
								}
							}
						}
					}
					result += max
				}
				dp[row][col1][col2] = result
			}
		}
	}
	return dp[0][0][n-1]
}

func cherryPicku_Another(grid [][]int) int {
	R := len(grid)
	C := len(grid[0])
	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	const _INF int = -1e9

	memoise := map[[3]int]int{}

	var solve func(int, int, int) int
	solve = func(r1, j1, j2 int) int {
		if r1 == R {
			return 0
		}
		if j1 < 0 || j2 < 0 || j2 >= C || j1 >= C {
			return _INF
		}

		if v, ok := memoise[[3]int{r1, j1, j2}]; ok {
			return v
		}

		var maxi int = 0
		for i := -1; i < 2; i++ {
			for j := -1; j < 2; j++ {
				if j1 == j2 {
					maxi = Max(maxi, grid[r1][j1]+solve(r1+1, j1+i, j2+j))
				} else {
					maxi = Max(maxi, grid[r1][j1]+grid[r1][j2]+solve(r1+1, j1+i, j2+j))
				}
			}
		}

		memoise[[3]int{r1, j1, j2}] = maxi
		return maxi
	}

	return Max(0, solve(0, 0, C-1))
}

func cherryPicku_2(grid [][]int) int {
	R := len(grid)
	C := len(grid[0])
	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	MaxV := func(args ...int) int {
		max := args[0]

		for i := 1; i < len(args); i++ {
			if max < args[i] {
				max = args[i]
			}
		}
		return max
	}

	const _INF int = -1e9

	memoise := map[[3]int]int{}

	var solve func(int, int, int) int
	solve = func(r1, j1, j2 int) int {
		if r1 >= R || j1 < 0 || j2 < 0 || j2 >= C || j1 >= C {
			return _INF
		}

		if v, ok := memoise[[3]int{r1, j1, j2}]; ok {
			return v
		}

		cherries := grid[r1][j1]
		if j1 != j2 {
			cherries += grid[r1][j2]
		}

		if r1 == R-1 {
			return cherries
		}

		p1 := solve(r1+1, j1-1, j2-1)
		p2 := solve(r1+1, j1, j2-1)
		p3 := solve(r1+1, j1+1, j2-1)

		p4 := solve(r1+1, j1-1, j2)
		p5 := solve(r1+1, j1, j2)
		p6 := solve(r1+1, j1+1, j2)

		p7 := solve(r1+1, j1-1, j2+1)
		p8 := solve(r1+1, j1, j2+1)
		p9 := solve(r1+1, j1+1, j2+1)

		cherries += MaxV(p1, p2, p3, p4, p5, p6, p7, p8, p9)
		memoise[[3]int{r1, j1, j2}] = cherries
		return cherries
	}

	return Max(0, solve(0, 0, C-1))
}
