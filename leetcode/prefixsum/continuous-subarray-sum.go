package prefixsum

// https://leetcode.com/problems/continuous-subarray-sum
func CheckSubarraySum(nums []int, k int) bool {
	idxMap := map[int]int{
		0: -1,
	}
	prefix := 0
	var remainder int
	for i, n := range nums {
		prefix += n
		remainder = prefix % k
		if _, ok := idxMap[remainder]; !ok {
			idxMap[remainder] = i
		} else if i-idxMap[remainder] > 1 {
			return true
		}
	}
	return false
}
