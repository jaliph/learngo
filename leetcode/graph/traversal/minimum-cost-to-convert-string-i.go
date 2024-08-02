package graph

// https://leetcode.com/problems/minimum-cost-to-convert-string-i/
func MinimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	const INF int = 1e9
	R := 27
	C := 27
	dist := make([][]int, R)
	rows := make([]int, R*C)
	for i := range rows {
		rows[i] = INF
	}
	for i := range dist {
		dist[i] = rows[i*C : (i+1)*C]
	}

	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	n := len(original)
	for i := 0; i < n; i++ {
		if cost[i] < dist[original[i]-'a'][changed[i]-'a'] {
			dist[original[i]-'a'][changed[i]-'a'] = cost[i]
		}
	}

	for i := 0; i < R; i++ {
		dist[i][i] = 0
	}

	for k := 0; k < R; k++ {
		for i := 0; i < R; i++ {
			for j := 0; j < R; j++ {
				if i == j {
					continue
				}
				dist[i][j] = Min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}

	var res int64 = 0
	for i := 0; i < len(source); i++ {
		s := source[i] - 'a'
		t := target[i] - 'a'
		// fmt.Println(dist[s][t])
		if dist[s][t] == INF {
			return -1
		} else {
			res += int64(dist[s][t])
		}
	}

	return res
}
