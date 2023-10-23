package twopointer

// https://leetcode.com/problems/maximum-score-of-a-good-subarray/

func MaximumScore(nums []int, k int) int {
	res := nums[k]
	l, r := k, k
	currMin := nums[k]

	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for l > 0 && r < len(nums)-1 {
		var left int
		var right int
		if l > 0 {
			left = nums[l-1]
		} else {
			left = 0
		}

		if r < len(nums)-1 {
			right = nums[r+1]
		} else {
			right = 0
		}

		if left < right {
			currMin = Min(currMin, right)
			r++
		} else {
			currMin = Min(currMin, left)
			l--
		}
		res = Max(res, currMin*(r-l+1))
	}
	return res
}
