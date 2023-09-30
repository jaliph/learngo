// https://leetcode.com/problems/remove-element/submissions/
package basic

func RemoveElement(nums []int, val int) int {
	w := 0

	for r := 0; r < len(nums); r++ {
		if nums[r] == val {
			continue
		}
		nums[w] = nums[r]
		w++
	}

	return w
}
