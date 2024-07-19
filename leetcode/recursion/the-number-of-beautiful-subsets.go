package recursion

import "sort"

// https://leetcode.com/problems/the-number-of-beautiful-subsets/
func BeautifulSubsets(nums []int, k int) int {

	count := 0
	arr := []int{}

	Abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	var dfs func(int, *[]int)
	dfs = func(idx int, temp *[]int) {
		if len(*temp) > 0 {
			count += 1
		}

		for i := idx; i < len(nums); i++ {
			// take
			canTake := true
			for j := 0; j < len(*temp); j++ {
				if Abs(nums[i]-(*temp)[j]) == k {
					canTake = false
					break
				}
			}

			if canTake {
				*temp = append(*temp, nums[i])
				dfs(i+1, temp)
				*temp = (*temp)[:len(*temp)-1]
			}
		}
	}

	sort.Ints(nums)
	dfs(0, &arr)

	return count
}
