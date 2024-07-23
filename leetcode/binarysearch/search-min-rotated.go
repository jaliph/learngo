package binarysearch

// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
func FindMin(nums []int) int {
	Min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}
	l := 0
	r := len(nums) - 1
	res := 10000
	var mid int
	for l <= r {
		if nums[l] <= nums[r] {
			res = Min(res, nums[l])
			break
		}

		mid = l + ((r - l) / 2)
		// left side is sorted
		if nums[l] <= nums[mid] {
			res = Min(res, nums[l])
			l = mid + 1
		} else { // If right part is sorted:
			res = Min(res, nums[mid])
			r = mid - 1
		}
	}

	return res
}
