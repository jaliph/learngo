package greedy

import (
	"sort"
)

// minimum-increment-to-make-array-unique

// TLE
func MinIncrementForUnique_TLE(nums []int) int {
	set := map[int]struct{}{}

	sort.Ints(nums)

	sum := 0
	for _, v := range nums {
		if _, ok := set[v]; ok {
			temp := v + 1
			_, ok := set[temp]
			for ok {
				temp = temp + 1
				_, ok = set[temp]
			}
			sum += temp - v
			set[temp] = struct{}{}
		} else {
			set[v] = struct{}{}
		}
	}
	return sum
}

func MinIncrementForUnique(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i := 1; i < len(nums); i++ {
		// if previous is greater or equal
		if nums[i-1] >= nums[i] {
			sum += (nums[i-1] - nums[i] + 1) // cost is diff + 1
			nums[i] = nums[i-1] + 1          // curr is prev + 1
		}
	}
	return sum
}
