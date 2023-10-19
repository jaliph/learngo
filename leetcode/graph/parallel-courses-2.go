package graph

func minNumberOfSemesters(n int, relations [][]int, k int) int {

	q := []int{}
	g := make([][]int, n+1)
	indegree := make([]int, n+1)
	dist := make([]int, n+1)
	countMap := map[int]int{}

	for _, v := range relations {
		if len(g[v[0]]) == 0 {
			g[v[0]] = []int{}
		}
		indegree[v[1]]++
		g[v[0]] = append(g[v[0]], v[1])
	}

	for i := 1; i <= n; i++ {
		if indegree[i] == 0 {
			dist[i] = 1
			countMap[dist[i]]++
			q = append(q, i)
		}
	}

	ans := 0
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		for _, neigh := range g[curr] {
			indegree[neigh]--
			dist[neigh] = dist[curr] + 1
			if indegree[neigh] == 0 {
				countMap[dist[neigh]]++
				q = append(q, neigh)
			}
		}
	}

	for _, v := range countMap {
		ans += (v / k)
		if v%k > 0 {
			ans++
		}
	}

	return ans
}

// func Driver() {
// 	n := 7
// 	relations := [][]int{{2, 1}, {3, 1}, {4, 1}, {1, 5}, {1, 6}, {7, 1}}
// 	k := 2

// 	fmt.Println(minNumberOfSemesters(n, relations, k))
// }
