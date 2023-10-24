// https://www.lintcode.com/problem/3673/

package graph

import (
	"fmt"
)

/**
 * @param n: the number of courses
 * @param relations: the relationship between all courses
 * @return: ehe minimum number of semesters required to complete all courses
 */
func MinimumSemesters(n int, relations [][]int) int {
	// write your code here

	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	g := make([][]int, n)
	indegree := make([]int, n)

	for _, v := range relations {
		if len(g[v[0]-1]) == 0 {
			g[v[0]-1] = []int{}
		}
		indegree[v[1]-1]++
		g[v[0]-1] = append(g[v[0]-1], v[1]-1)
	}

	q := [][]int{}
	for i, v := range indegree {
		if v == 0 {
			q = append(q, []int{i, 1})
		}
	}

	sems := 0

	for len(q) > 0 {

		curr := q[0]
		fmt.Println(curr)
		q = q[1:]
		sems = Max(sems, curr[1])

		for _, neigh := range g[curr[0]] {
			indegree[neigh]--

			if indegree[neigh] == 0 {
				q = append(q, []int{neigh, curr[1] + 1})
			}
		}
	}

	if len(q) == 0 && sems != 0 {
		return sems
	} else {
		return -1
	}
}

func Driver() {
	n := 4
	relations := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}}
	fmt.Println(MinimumSemesters(n, relations))
}
