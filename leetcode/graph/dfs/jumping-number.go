package graph

// fmt.Println(graph.JumpingNumbers(20))
func JumpingNumbers(n int) []int {
	res := []int{}

	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var dfs func(i int)
	dfs = func(i int) {
		if i > n {
			return
		}

		if i <= n {
			res = append(res, i)
		}

		ld := i % 10
		if ld == 0 {
			dfs((i * 10) + 1)
		} else if ld == 9 {
			dfs((i * 10) + 8)
		} else {
			dfs((i * 10) + (ld + 1))
			dfs((i * 10) + (ld - 1))
		}
	}
	for i := 1; i <= Min(9, n); i++ {
		dfs(i)
	}

	return res
}
