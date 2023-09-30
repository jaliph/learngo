package stacks

func Find132pattern(nums []int) bool {
	s := [][]int{}

	currMin := nums[0]

	for i := 1; i < len(nums); i++ {
		for len(s) > 0 && s[len(s)-1][0] <= nums[i] {
			s = s[:len(s)-1]
		}

		if len(s) > 0 && s[len(s)-1][1] < nums[i] {
			return true
		}

		s = append(s, []int{nums[i], currMin})
		currMin = MinInt(currMin, nums[i])
	}

	return false
}

func MinInt(args ...int) int {
	min := args[0]

	for _, v := range args {
		if min > v {
			min = v
		}
	}
	return min
}
