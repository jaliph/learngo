package dp

func numberWays(hats [][]int) int {
	const M = 1e9 + 7
	people := map[int][]int{}

	for i := 0; i < len(hats); i++ {
		for j := 0; j < len(hats[i]); j++ {
			if _, ok := people[hats[i][j]]; !ok {
				people[hats[i][j]] = []int{}
			}
			people[hats[i][j]] = append(people[hats[i][j]], i)
		}
	}

	p := [][]int{}
	for _, v := range people {
		p = append(p, v)
	}

	var solve func(int, int) int
	solve = func(idx, mask int) int {
		// base
		if mask == 1<<len(hats)-1 {
			return 1
		}
		if idx == len(p) {
			return 0
		}

		// recur
		ans := 0

		for _, pep := range p[idx] {
			if (mask >> pep & 1) == 0 {
				ans += solve(idx+1, mask|1<<pep)
				ans = ans % M
			}
		}

		ans += solve(idx+1, mask)
		ans = ans % M
		return ans
	}

	return solve(0, 0)
}
