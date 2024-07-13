package graph

import (
	"fmt"
	"sort"
)

func GetAncestors(n int, edges [][]int) [][]int {
	g := map[int][]int{}
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
	}

	res := make([][]int, n)
	var visited []bool

	var dfs func(curr int, par int)
	dfs = func(curr int, par int) {
		fmt.Println(curr)
		visited[curr] = true
		for _, neigh := range g[curr] {
			if !visited[neigh] {
				res[neigh] = append(res[neigh], par)
				dfs(neigh, par)
			}
		}
	}

	for i := 0; i < n; i++ {
		visited = make([]bool, n)
		dfs(i, i)
	}

	for _, arr := range res {
		sort.Ints(arr)
	}

	return res
}
