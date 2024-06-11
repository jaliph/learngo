package basic

import "sort"

// https://leetcode.com/problems/rank-transform-of-an-array/
func arrayRankTransform(arr []int) []int {
	idxMap := map[int]int{}

	for i, v := range arr {
		idxMap[v] = i
	}

	indexes := []int{}
	for k := range idxMap {
		indexes = append(indexes, k)
	}
	sort.Slice(indexes, func(i, j int) bool {
		return indexes[i] < indexes[j]
	})

	for i, v := range indexes {
		idxMap[v] = i + 1
	}

	for i, v := range arr {
		arr[i] = idxMap[v]
	}

	return arr
}
