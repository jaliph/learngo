package dp

import (
	"math"
)

func jump(nums []int) int {
	dest := len(nums) - 1

	dp := make(map[int]int)
	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

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
				ans = Min(ans, subProb+1)
			}
		}

		dp[i] = ans
		return ans
	}

	return solve(0)
}
