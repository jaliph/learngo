package cyclic

func FindMex(nums []int) int {
	for i := 0; i < len(nums); i++ {
		val := abs(nums[i])
		if 0 <= val && val <= len(nums)-1 {
			if nums[val] > 0 {
				nums[val] *= -1
			}
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return i
		}
	}
	return len(nums)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
