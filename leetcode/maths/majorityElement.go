package maths

// https://leetcode.com/problems/majority-element
func MajorityElement(nums []int) int {
	cnt := 0
	res := 0

	for _, i := range nums {
		if cnt == 0 {
			res = i
		}

		if res == i {
			cnt++
		} else {
			cnt--
		}
	}
	return res
}

func MajorityElementBrute(nums []int) int {
	cnt := 0
	res := 0

	cMap := make(map[int]int)

	for _, i := range nums {
		cMap[i]++
		if cnt < cMap[i] {
			res = i
			cnt = cMap[i]
		}
	}
	return res
}
