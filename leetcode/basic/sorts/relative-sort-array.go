package basic

import "sort"

func RelativeSortArray(arr1 []int, arr2 []int) []int {
	n1Map := map[int]int{}
	for i, n := range arr2 {
		n1Map[n] = i
	}

	MAX := (1 << 9) + 7

	sort.Slice(arr1, func(i, j int) bool {
		var x, y int
		if val, ok := n1Map[arr1[i]]; ok {
			x = val
		} else {
			x = MAX
		}

		if val, ok := n1Map[arr1[j]]; ok {
			y = val
		} else {
			y = MAX
		}

		return x < y
	})

	return arr1
}
