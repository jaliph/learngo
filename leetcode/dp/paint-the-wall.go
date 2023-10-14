package dp

func paintWalls(cost []int, time []int) int {

	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	dp := map[[2]int]int{}

	var solve func(int, int) int
	solve = func(idx int, remaining int) int {
		// base
		if remaining <= 0 {
			return 0
		}

		if idx == len(time) {
			return 1e9
		}

		if v, ok := dp[[2]int{idx, remaining}]; ok {
			return v
		}

		// recur
		paint := cost[idx] + solve(idx+1, remaining-1-time[idx])
		skip := solve(idx+1, remaining)

		dp[[2]int{idx, remaining}] = Min(paint, skip)
		return dp[[2]int{idx, remaining}]
	}

	return solve(0, len(cost))
}
