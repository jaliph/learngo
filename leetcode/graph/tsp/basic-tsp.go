package graph

import "fmt"

const INVALID int = 1e9

func TSP(g [][]int) int {

	ans := 0
	visited := map[int]bool{}
	path := []int{}
	var solve func(int)
	solve = func(curr int) {

		nextIdx := INVALID
		cost := INVALID
		visited[curr] = true
		path = append(path, curr)

		for i, v := range g[curr] {
			_, ok1 := visited[i]
			if !ok1 && v > 0 && v < cost {
				cost = v
				nextIdx = i
			}
		}

		if cost != INVALID {
			ans += cost
		}

		if nextIdx == INVALID {
			nextIdx = 0
			ans += g[curr][nextIdx]
			return
		}
		solve(nextIdx)
	}

	solve(0)
	fmt.Println(path)
	return ans
}

func TSP_1(g [][]int) int {
	source := 0
	const INF int = 1e9

	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}
	N := len(g)
	var solve func(int, int) int
	solve = func(curr, mask int) int {
		// base
		if (1<<N)-1 == mask {
			return g[curr][source]
		}
		// recur

		ans := INF
		for i, v := range g[curr] {
			if (mask>>i)&1 == 0 {
				subans := v + solve(i, mask^(1<<i))
				ans = Min(ans, subans)
			}
		}
		return ans
	}
	return solve(source, 1)
}
