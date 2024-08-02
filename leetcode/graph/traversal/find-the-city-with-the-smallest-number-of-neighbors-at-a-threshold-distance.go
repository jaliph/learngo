package graph

// https://leetcode.com/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance
func FindTheCity(n int, edges [][]int, distanceThreshold int) int {
	const INF int = 1e9
	R := n
	C := n
	Min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}

	dist := make([][]int, R)
	rows := make([]int, R*C)
	for i := range rows {
		rows[i] = INF
	}
	for i := range dist {
		dist[i] = rows[i*C : (i+1)*C]
	}

	for _, e := range edges {
		dist[e[0]][e[1]] = e[2]
		dist[e[1]][e[0]] = e[2]
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i == j {
					continue
				}
				dist[i][j] = Min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}

	res := 0
	city := 0
	distList := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if dist[i][j] <= distanceThreshold {
				distList[i]++
			}
		}
		if res >= distList[i] {
			res = distList[i]
			city = i
		}
	}
	return city
}
