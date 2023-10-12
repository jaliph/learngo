package dp

import (
	"fmt"
	"math"
)

func jump(nums []int) int {
	dest := len(nums) - 1

	dp := make(map[int]int)

	var solve func(i int) int
	solve = func(i int) int {
		if i == dest {
			return 0
		}
		if i > dest {
			return math.MaxInt
		}

		if v, ok := dp[i]; ok {
			return v
		}

		ans := math.MaxInt
		jumps := nums[i]
		for j := 1; j <= jumps; j++ {
			subProb := solve(j + i)
			if subProb != math.MaxInt {
				ans = MinInt(ans, subProb+1)
			}
		}

		dp[i] = ans
		return ans
	}

	return solve(0)
}

func MinInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Driver() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}
