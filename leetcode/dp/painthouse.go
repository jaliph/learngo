package dp

import "fmt"

/**
 * @param costs: n x 3 cost matrix
 * @return: An integer, the minimum cost to paint all houses
 */
func MinCost(costs [][]int) int {
	// write your code here

	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}
	prev := costs[0]
	for i := 1; i < len(costs); i++ {
		dp := make([]int, 3)
		for j := 0; j < 3; j++ {
			var temp int = 1e9
			for _j := 0; _j < 3; _j++ {
				if _j != j {
					temp = Min(temp, prev[_j])
				}
			}
			dp[j] = costs[i][j] + temp
		}

		prev = dp
	}

	return Min(prev[0], Min(prev[1], prev[2]))
}

func Driver() {
	fmt.Println(MinCost([][]int{{1, 2, 3}, {1, 4, 6}}))
}
