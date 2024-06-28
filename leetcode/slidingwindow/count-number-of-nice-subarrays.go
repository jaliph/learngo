package slidingwindow

// https://leetcode.com/problems/count-number-of-nice-subarrays
// nums = [1,1,2,1,1], k = 3
func NumberOfSubarrays(nums []int, k int) int {

	wStart := 0
	count := 0
	oddCount := 0
	res := 0
	for wEnd := 0; wEnd < len(nums); wEnd++ {
		if nums[wEnd]&1 != 0 {
			oddCount++
			count = 0
		}

		for oddCount == k {
			count++
			if nums[wStart]&1 != 0 {
				oddCount--
			}
			wStart++
		}
		// use the count for future subarray until another odd is found
		res += count
	}

	return res
}
