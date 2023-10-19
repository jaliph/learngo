// https://leetcode.com/problems/shortest-path-visiting-all-nodes/description/

package graph

func shortestPathLength(graph [][]int) int {

	q := [][3]int{}

	n := len(graph)

	visited := map[[2]int]bool{}
	for i := range graph {
		visited[[2]int{i, 1 << i}] = true
		q = append(q, [3]int{i, 1 << i, 0})
	}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if curr[1] == (1<<n)-1 {
			return curr[2]
		}

		for _, neigh := range graph[curr[0]] {
			newMask := curr[1] | (1 << neigh)
			if _, ok := visited[[2]int{neigh, newMask}]; !ok {
				visited[[2]int{neigh, newMask}] = true
				q = append(q, [3]int{neigh, newMask, curr[2] + 1})
			}
		}
	}
	return 0
}
