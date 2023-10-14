package dp

import (
	"sort"
)

func maxHeight(cuboids [][]int) int {

	for _, v := range cuboids {
		sort.Ints(v)
	}

	cuboids = append(cuboids, []int{0, 0, 0})

	sort.Slice(cuboids, func(i, j int) bool {
		if cuboids[i][0] != cuboids[j][0] {
			return cuboids[i][0] < cuboids[j][0]
		} else if cuboids[i][1] != cuboids[j][1] {
			return cuboids[i][1] < cuboids[j][1]
		} else {
			return cuboids[i][2] < cuboids[j][2]
		}
	})

	dp := make([]int, len(cuboids))
	dp[0] = cuboids[0][2]

	MaxInt := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	max := 0
	for i := 1; i < len(dp); i++ {
		dp[i] = cuboids[i][2]
		for j := 0; j < i; j++ {
			if cuboids[j][0] <= cuboids[i][0] && cuboids[j][1] <= cuboids[i][1] && cuboids[j][2] <= cuboids[i][2] {
				dp[i] = MaxInt(dp[i], dp[j]+cuboids[i][2])
				max = MaxInt(max, dp[i])
			}
		}
	}

	return max
}
