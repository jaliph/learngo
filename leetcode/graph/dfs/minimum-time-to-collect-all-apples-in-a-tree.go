package graph

// https://leetcode.com/problems/minimum-time-to-collect-all-apples-in-a-tree/

func minTime(n int, edges [][]int, hasApple []bool) int {

	g := map[int][]int{}
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}

	var dfs func(node, par int) int
	dfs = func(node, par int) int {
		time := 0
		for _, neigh := range g[node] {
			if neigh != par {
				subSolve := dfs(neigh, node)
				if subSolve > 0 || hasApple[neigh] {
					time += subSolve + 2
				}
			}
		}
		return time
	}

	return dfs(0, -1)
}
