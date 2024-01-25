package dp

func rob(nums []int) int {

	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	if len(nums) == 1 {
		return nums[0]
	}

	prev1 := 0 // dp[i - 1]
	prev2 := 0 // dp[i - 2]
	for i := 0; i < len(nums); i++ {
			 := Max(prev1, prev2+nums[i])
		prev2 = prev1
		prev1 = curr
	}
	return prev1
}
