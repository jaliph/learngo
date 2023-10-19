package dp

func Rob(nums []int) int {
	Max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	helper := func(houses []int) int {
		prev1, prev2, curr := 0, 0, 0
		for _, h := range houses {
			curr = Max(prev1, prev2+h)
			prev2 = prev1
			prev1 = curr
		}
		return prev1
	}

	return Max(nums[0], Max(helper(nums[1:]), helper(nums[:len(nums)-1])))
}
