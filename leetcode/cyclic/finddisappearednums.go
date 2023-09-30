package cyclic

// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array
func FindDisappearedNumbers(nums []int) []int {
	i := 0
	for i < len(nums) {
		j := nums[i] - 1
		if nums[i] != nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			i++
		}
		// fmt.Println(nums)
	}

	res := []int{}

	for i, v := range nums {
		if v != i+1 {
			res = append(res, i+1)
		}
	}
	return res
}
