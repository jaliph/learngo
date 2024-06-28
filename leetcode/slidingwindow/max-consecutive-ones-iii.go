package slidingwindow

// https://leetcode.com/problems/max-consecutive-ones-iii/
func longestOnes(nums []int, k int) int {

	wStart := 0
	countZero := 0
	res := 0
	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for wEnd := 0; wEnd < len(nums); wEnd++ {
		if nums[wEnd] == 0 {
			countZero++
		}

		// cannot do more than k skips.. so lets shrink the window
		for countZero > k {
			if nums[wStart] == 0 {
				countZero--
			}
			wStart++
		}
		res = Max(res, wEnd-wStart+1)
	}
	return res
}
