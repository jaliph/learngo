// https://leetcode.com/problems/parallel-courses-iii/

package graph

func MinimumTime_Graph(n int, relations [][]int, time []int) int {

	q := []int{}

	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	g := make([][]int, n+1)
	indegree := make([]int, n+1)
	waitTime := make([]int, n+1)
	maxWaitTime := 0

	for _, v := range relations {
		if len(g[v[0]]) == 0 {
			g[v[0]] = []int{}
		}
		indegree[v[1]]++
		g[v[0]] = append(g[v[0]], v[1])
	}

	for i := 1; i <= n; i++ {
		if indegree[i] == 0 {
			waitTime[i] = time[i-1]
			maxWaitTime = Max(maxWaitTime, waitTime[i])
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		for _, child := range g[curr] {
			indegree[child]--
			waitTime[child] = Max(waitTime[child], time[child-1]+waitTime[curr])
			maxWaitTime = Max(maxWaitTime, waitTime[child])
			if indegree[child] == 0 {
				q = append(q, child)
			}
		}
	}
	// fmt.Println(waitTime)

	return maxWaitTime
}

func minimumTime_DP(n int, relations [][]int, time []int) int {

	g := make([][]int, n)

	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for _, v := range relations {
		if len(g[v[0]-1]) == 0 {
			g[v[0]-1] = []int{}
		}
		g[v[0]-1] = append(g[v[0]-1], v[1]-1)
	}

	memoise := map[int]int{}
	max := 0

	var solve func(int) int
	solve = func(i int) int {

		if v, ok := memoise[i]; ok {
			return v
		}

		waitTime := 0
		for _, v := range g[i] {
			waitTime = Max(waitTime, solve(v))
		}

		waitTime = time[i] + waitTime
		memoise[i] = waitTime
		return waitTime
	}

	for i := 0; i < n; i++ {
		max = Max(max, solve(i))
	}

	return max
}

// func Driver() {
// 	n := 5
// 	relations := [][]int{{1, 5}, {2, 5}, {3, 5}, {3, 4}, {4, 5}}
// 	time := []int{1, 2, 3, 4, 5}

// 	fmt.Println(minimumTime(n, relations, time))

// 	fmt.Println(minimumTime(2, [][]int{{2, 1}}, []int{100, 200}))
// }
