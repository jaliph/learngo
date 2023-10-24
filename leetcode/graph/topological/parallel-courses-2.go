package graph

import "fmt"

func minNumberOfSemesters(n int, relations [][]int, k int) int {

	q := [][2]int{}
	g := make([][]int, n)
	indegree := make([]int, n)

	// form the graph
	for _, v := range relations {
		if len(g[v[0]-1]) == 0 {
			g[v[0]-1] = []int{}
		}
		indegree[v[1]-1]++
		g[v[0]-1] = append(g[v[0]-1], v[1]-1)
	}

	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// calculating depth
	depth := make([]int, n)
	for i := range depth {
		depth[i] = -1
	}

	var depthFinder func(int) int
	depthFinder = func(i int) int {
		if depth[i] != -1 {
			return depth[i]
		}

		d := 0
		for _, v := range g[i] {
			d = Max(d, depthFinder(v))
		}
		depth[i] = d + 1
		return d + 1
	}

	for i := 0; i < n; i++ {
		depthFinder(i)
	}

	fmt.Println(depth)

	// Priority Queue
	BinaryInsert := func(node [2]int, arr *[][2]int) {
		l, r := 0, len(*arr)-1

		var mid int
		for l <= r {
			mid = l + (r-l)/2
			if node[1] == (*arr)[mid][1] {
				l = mid
				break
			} else if node[1] > (*arr)[mid][1] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		if l == len(*arr) {
			(*arr) = append((*arr), node)
		} else {
			(*arr) = append((*arr)[:l+1], (*arr)[l:]...)
			(*arr)[l] = node
		}
	}

	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			BinaryInsert([2]int{i, depth[i]}, &q)
		}
	}

	ans := 0
	for len(q) > 0 {
		fmt.Println("q", q)

		ans++
		temp := [][2]int{}
		for i, j := 0, len(q); i < k && j > 0; i, j = i+1, j-1 {
			curr := q[0]
			q = q[1:]
			fmt.Println("Traversing", curr[0]+1)
			for _, v := range g[curr[0]] {
				indegree[v]--

				if indegree[v] == 0 {
					temp = append(temp, [2]int{v, depth[v]})
				}
			}
		}
		fmt.Println(temp)
		for _, v := range temp {
			BinaryInsert(v, &q)
		}
	}

	return ans
}

// func Driver() {
// 	n := 12
// 	relations := [][]int{{11, 10}, {6, 3}, {2, 5}, {9, 2}, {4, 12}, {8, 7}, {9, 5}, {6, 2}, {7, 2}, {7, 4}, {9, 3}, {11, 1}, {4, 3}}
// 	k := 3

// 	fmt.Println(minNumberOfSemesters(n, relations, k))
// }
