package graph

// https://leetcode.com/problems/minimum-falling-path-sum/

func minFallingPathSum(matrix [][]int) int {

	var INF int = 1e9
	R := len(matrix)
	C := len(matrix[0])

	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}
	visited := map[[2]int]int{}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {

		if i == R-1 {
			return matrix[i][j]
		}

		if i >= R || j < 0 || j >= C {
			return INF
		}

		if v, ok := visited[[2]int{i, j}]; ok {
			return v
		}

		step := matrix[i][j]
		res := INF

		res = Min(res, step+dfs(i+1, j))
		if j-1 >= 0 {
			res = Min(res, step+dfs(i+1, j-1))
		}

		if j+1 < C {
			res = Min(res, step+dfs(i+1, j+1))
		}

		// if _, ok := visited[[2]int{i, j}]; ok {
		// 	return
		// }
		visited[[2]int{i, j}] = res

		return res
	}

	ans := INF
	for i := 0; i < C; i++ {
		ans = Min(ans, dfs(0, i))
	}

	return ans
}

func minFallingPathSumBrute(matrix [][]int) int {

	var res int = 1e9
	R := len(matrix)
	C := len(matrix[0])

	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}
	// visited := map[[2]int]bool{}

	var dfs func(i, j, sum int)
	dfs = func(i, j, sum int) {

		if i == R-1 {
			// fmt.Println(path, sum)
			res = Min(res, sum)
			return
		}

		if j < 0 || j >= C {
			return
		}

		// if _, ok := visited[[2]int{i, j}]; ok {
		// 	return
		// }
		// visited[[2]int{i, j}] = true

		dfs(i+1, j, sum+matrix[i][j])
		dfs(i+1, j-1, sum+matrix[i][j])
		dfs(i+1, j+1, sum+matrix[i][j])

		// delete(visited, [2]int{i, j})
	}

	for i := 0; i < C; i++ {
		dfs(0, i, 0)
	}

	return res
}

// func Driver() {
// 	mat := [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}
// 	fmt.Println(minFallingPathSum(mat))
// }
