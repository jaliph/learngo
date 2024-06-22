package prefixsum

// https://leetcode.com/problems/count-number-of-nice-subarrays
// nums = [1,1,2,1,1], k = 3
func NumberOfSubarrays(nums []int, k int) int {

	prefixCounts := map[int]int{0: 1}
	oddCount := 0
	count := 0

	for _, num := range nums {
		if num%2 != 0 {
			oddCount++
		}
		if v, ok := prefixCounts[oddCount-k]; ok {
			count += v
		}
		prefixCounts[oddCount]++
	}

	return count
}
