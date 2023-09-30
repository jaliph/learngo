// https://leetcode.com/problems/contains-duplicate/

package basic

func ContainsDuplicate(nums []int) bool {
	set := make(map[int]bool)

	for _, v := range nums {
		if _, ok := set[v]; ok {
			return true
		}
		set[v] = true
	}

	return false
}
