package slidingwindow

import "sort"

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)

	wStart := 0
	wSum := 0

	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	maxFrequency := 1
	for wEnd := 0; wEnd < len(nums); wEnd++ {
		wSum += nums[wEnd]

		for (nums[wEnd]*(wEnd-wStart+1))-wSum > k {
			wSum -= nums[wStart]
			wStart++
		}

		maxFrequency = Max(maxFrequency, wEnd-wStart+1)
	}

	return maxFrequency
}
