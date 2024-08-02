package slidingwindow

// https://leetcode.com/problems/minimum-swaps-to-group-all-1s-together-ii
func MinSwaps(nums []int) int {

	total_ones := 0
	for i := range nums {
		total_ones += nums[i]
	}
	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// in the sliding window of size total ones, count the max ones we can achieve. swaps is diff of that
	// for circular part, append nums with itself, to do this % n after exceeding n

	max_ones := 0
	win_ones := 0
	wStart := 0
	n := len(nums)
	for wEnd := 0; wEnd < 2*n; wEnd++ {

		win_ones += nums[wEnd%n]

		if wEnd-wStart+1 > total_ones {
			win_ones -= nums[wStart%n]
			wStart++
		}

		max_ones = Max(max_ones, win_ones)
	}

	return total_ones - max_ones
}
