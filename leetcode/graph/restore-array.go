package graph

// https://leetcode.com/problems/restore-the-array-from-adjacent-pairs/

func RestoreArray(adjacentPairs [][]int) []int {
	res := []int{}

	g := map[int][]int{}
	for _, pairs := range adjacentPairs {
		if _, ok := g[pairs[0]]; !ok {
			g[pairs[0]] = []int{}
		}
		if _, ok := g[pairs[1]]; !ok {
			g[pairs[1]] = []int{}
		}
		g[pairs[0]] = append(g[pairs[0]], pairs[1])
		g[pairs[1]] = append(g[pairs[1]], pairs[0])
	}

	// find a node which has only 1 adjust
	for i, v := range g {
		if len(v) == 1 {
			res = append(res, i, v[0])
			break
		}
	}

	for len(res) < len(adjacentPairs)+1 {
		last := res[len(res)-1]
		prev := res[len(res)-2]

		candidates := g[last]

		if prev != candidates[0] {
			res = append(res, candidates[0])
		} else {
			res = append(res, candidates[1])
		}
	}

	return res
}

// func Driver() {
// 	pairs := [][]int{{2, 1}, {3, 4}, {3, 2}}
// 	fmt.Println(restoreArray(pairs))
// }
