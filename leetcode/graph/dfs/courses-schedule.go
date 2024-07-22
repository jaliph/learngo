package graph

import "fmt"

func CanFinish(numCourses int, prerequisites [][]int) bool {

	g := map[int][]int{}
	for _, reqs := range prerequisites {
		g[reqs[0]] = append(g[reqs[0]], reqs[1])
	}

	visited := make([]int, numCourses)

	var dfs func(curr, par int) bool
	dfs = func(curr, par int) bool {
		visited[curr] = 1

		ans := true
		for _, neigh := range g[curr] {
			if visited[neigh] == 0 {
				ans = ans && dfs(neigh, curr)
			} else if visited[neigh] == 1 {
				return false
			}
		}
		visited[curr] = 2
		return ans
	}

	for i := 0; i < numCourses; i++ {
		if visited[i] == 0 {
			fmt.Println(dfs(i, -1))
			if !dfs(i, -1) {
				return false
			}
		}
	}

	return true
}
