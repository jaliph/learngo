package graph

import "fmt"

// https://leetcode.com/problems/build-a-matrix-with-conditions
func sortDeps(conds [][]int, k int) []int {
	g := map[int][]int{}
	for _, c := range conds {
		g[c[0]] = append(g[c[0]], c[1])
	}

	visited := make([]int, k+1)
	order := []int{}
	var dfs func(int) bool
	dfs = func(curr int) bool {
		visited[curr] = 1
		ans := false
		for _, neigh := range g[curr] {
			if visited[neigh] == 0 {
				ans = ans || dfs(neigh)
			} else if visited[neigh] == 1 {
				return true
			}
		}
		visited[curr] = 2
		order = append(order, curr)
		return ans
	}

	for i := 1; i <= k; i++ {
		if visited[i] == 0 {
			if dfs(i) {
				return []int{}
			}
		}
	}

	return func() []int {
		l := 0
		r := len(order) - 1
		for l <= r {
			order[l], order[r] = order[r], order[l]
			l++
			r--
		}
		return order
	}()
}

func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {

	rowOrder := sortDeps(rowConditions, k)
	colOrder := sortDeps(colConditions, k)
	fmt.Println(rowOrder, colOrder)

	if len(rowOrder) == 0 || len(colOrder) == 0 {
		return [][]int{}
	}

	rowMap := map[int]int{}
	colMap := map[int]int{}

	for idx, r := range rowOrder {
		rowMap[r] = idx
	}

	for idx, c := range colOrder {
		colMap[c] = idx
	}

	res := make([][]int, k)
	for i := range res {
		res[i] = make([]int, k)
	}

	for i := 1; i <= k; i++ {
		rIdx := rowMap[i]
		cIdx := colMap[i]
		res[rIdx][cIdx] = i
	}
	return res
}
