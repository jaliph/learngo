// https://leetcode.com/problems/minimum-number-of-operations-to-make-array-continuous/
package binarysearch

import (
	"fmt"
	"math"
	"sort"
)

func MinOperations(nums []int) int {
	set := make(map[int]bool)
	l := len(nums)

	uniq := []int{}

	for _, i := range nums {
		if _, ok := set[i]; !ok {
			uniq = append(uniq, i)
		}
		set[i] = true
	}

	sort.SliceStable(uniq, func(i, j int) bool {
		return uniq[i] < uniq[j]
	})

	binarySearch := func(nums []int, target int) int {
		l, r := 0, len(nums)
		var mid int
		for l < r {
			mid = l + ((r - l) / 2)
			if target > nums[mid] {
				l = mid + 1
			} else {
				r = mid
			}
		}

		if l < len(nums) && nums[l] == target {
			return l + 1
		} else {
			return l
		}
	}

	ans := math.MaxInt
	fmt.Println(uniq)
	for i, v := range uniq {
		endElement := v + (l - 1)
		idx := binarySearch(uniq, endElement)
		ans = MinInt(ans, l-(idx-i))
	}
	return ans
}

func MinInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
