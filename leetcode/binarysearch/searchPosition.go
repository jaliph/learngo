package binarysearch

// https://leetcode.com/problems/search-insert-position
func SearchInsert(nums []int, target int) int {
	l := 0
	r := len(nums)

	var mid int
	for l < r {
		mid = l + ((r - l) / 2)
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}
