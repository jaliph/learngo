package graph

// https://leetcode.com/problems/cherry-pickup/

func cherryPickup(grid [][]int) int {
	R := len(grid)
	C := len(grid[0])

	var _INF int = -1e9

	memoise := map[[3]int]int{}

	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var solve func(int, int, int) int
	solve = func(r1, c1, c2 int) int {
		r2 := r1 + c1 - c2

		if r1 >= R || r2 >= R || c1 >= C || c2 >= C {
			return _INF
		}

		if grid[r1][c1] == -1 || grid[r2][c2] == -1 {
			return _INF
		}

		if v, ok := memoise[[3]int{r1, c1, c2}]; ok {
			return v
		}

		if r1 == R-1 && c1 == C-1 {
			return grid[r1][c1]
		}

		if r2 == R-1 && c2 == C-1 {
			return grid[r2][c2]
		}

		cherries := grid[r1][c1]
		if r1 != r2 || c1 != c2 {
			cherries += grid[r2][c2]
		}

		/** Paths
				A   B
				-----
				r   r
		    d   d
				r   d
				d   r
				**/
		p1 := solve(r1+1, c1, c2)
		p2 := solve(r1, c1+1, c2+1)
		p3 := solve(r1+1, c1, c2+1)
		p4 := solve(r1, c1+1, c2)

		cherries += Max(p1, Max(p2, Max(p3, p4)))
		memoise[[3]int{r1, c1, c2}] = cherries
		return cherries
	}

	return Max(0, solve(0, 0, 0))
}
