// https://leetcode.com/problems/monotonic-array/
package basic

func IsMonotonic(nums []int) bool {

	if len(nums) < 3 {
		return true
	}

	incr := false
	desc := false
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			incr = true
		}

		if nums[i-1] > nums[i] {
			desc = true
		}

		if incr && desc {
			return false
		}
	}

	return true
}
