// https://leetcode.com/problems/subarray-sum-equals-k
package prefixsum

func subarraySum(nums []int, k int) int {
	sumMap := map[int]int{}

	sumMap[0] = 1
	prefix := 0
	count := 0
	for _, n := range nums {
		prefix += n

		if v, ok := sumMap[prefix-k]; ok {
			count += v
		}

		sumMap[prefix]++
	}

	return count
}
