package twopointer

import (
	"fmt"
)

// https://leetcode.com/problems/count-nice-pairs-in-an-array/solutions/4311440/python3-two-sum-approach-number-of-pairs-n-n-1-2-beats-99/?envType=daily-question&envId=2023-11-21

func countNicePairs(nums []int) int {

	const MOD int = 1e9 + 7
	Rev := func(i int) int {
		res := 0

		for i > 0 {
			d := i % 10
			res = (res * 10) + d
			i /= 10
		}
		return res
	}

	res := 0
	countMap := map[int]int{}
	for _, v := range nums {
		key := v - Rev(v)
		res += countMap[key]
		res %= MOD
		countMap[key]++
	}
	return res % MOD
}

func Driver() {
	fmt.Println(countNicePairs([]int{13, 10, 35, 24, 76}))
}

func maxWidthOfVerticalArea(points [][]int) int {

	fmt.Println(points)
	return 0
}
