package dp

func canCross(stones []int) bool {

	if stones[1] != 1 {
		return false
	}

	dp := map[[2]int]int{}

	riverLn := len(stones)
	var solve func(int, int) int
	solve = func(idx, jump int) int {
		if idx == riverLn-1 {
			return 1
		}

		if v, ok := dp[[2]int{idx, jump}]; ok {
			return v
		}

		res := 0

		for j := idx + 1; j < riverLn; j++ {
			newJump := stones[j] - stones[idx]
			if newJump >= (jump-1) && newJump <= (jump+1) && solve(j, newJump) == 1 {
				res = 1
				break
			}
		}

		dp[[2]int{idx, jump}] = res
		return res
	}

	return solve(1, 1) == 1
}

func canCross_TD(stones []int) bool {

	stoneMap := make(map[int]int)

	for i, v := range stones {
		stoneMap[v] = i
	}

	dp := make(map[[2]int]int)

	riverLn := len(stones)
	var solve func(int, int) int

	solve = func(idx, jump int) int {
		if idx == riverLn-1 {
			return 1
		}

		if v, ok := dp[[2]int{idx, jump}]; ok {
			return v
		}

		if jump == 0 {
			return 0
		}

		stoneToJump := jump + stones[idx]

		var newIdx int = -1
		if v, ok := stoneMap[stoneToJump]; ok {
			newIdx = v
		} else {
			// not possible to jump to that stone
			dp[[2]int{idx, jump}] = 0
			return 0
		}

		if newIdx <= riverLn-1 {
			j1 := solve(newIdx, jump)
			j2 := solve(newIdx, jump+1)
			j3 := solve(newIdx, jump-1)
			res := (j1 | j2 | j3)
			dp[[2]int{idx, jump}] = res
			return res
		}
		dp[[2]int{idx, jump}] = 0
		return 0
	}

	return solve(0, 1) == 1
}

func canCrossBrute(stones []int) bool {
	// stones values to their indexes
	stoneMap := make(map[int]int)
	for i, v := range stones {
		stoneMap[v] = i
	}

	var dfs func(int, int) int

	dfs = func(i, jump int) int {
		if i == len(stones)-1 {
			return 1
		}

		// this will cause a cycle
		if jump == 0 {
			return 0
		}

		// possible jump
		stoneVal := jump + stones[i]
		var idx int

		if i, ok := stoneMap[stoneVal]; ok {
			idx = i
		} else {
			return 0
		}

		if idx <= len(stones)-1 {
			j1 := dfs(idx, jump)
			j2 := dfs(idx, jump-1)
			j3 := dfs(idx, jump+1)
			return (j1 | j2 | j3)
		}

		return 0
	}

	return dfs(0, 1) == 1
}
