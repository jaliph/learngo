package intervals

import "sort"

// https://leetcode.com/problems/merge-intervals/
func Merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}
	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	res := [][]int{}
	res = append(res, intervals[0])
	for i := range intervals {
		if i == 0 {
			continue
		}
		if intervals[i][0] > res[len(res)-1][1] {
			res = append(res, intervals[i])
		} else if res[len(res)-1][1] >= intervals[i][0] {
			res[len(res)-1][0] = Min(res[len(res)-1][0], intervals[i][0])
			res[len(res)-1][1] = Max(res[len(res)-1][1], intervals[i][1])
		}
	}
	return res
}
