package maths

import (
	"sort"
)

// https://leetcode.com/problems/binary-trees-with-factors/

func numFactoredBinaryTrees(arr []int) int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	const MOD int = 1e9 + 7

	trees := map[int]int{}

	// all nums are tree by themselves

	for _, v := range arr {
		trees[v] = 1
	}

	for _, root := range arr {
		for _, factor := range arr {
			if root == factor {
				break
			}

			// if factor represents one child then root/factor represent next.. total trees for root is their multiply

			_, ok := trees[root/factor]
			if root%factor == 0 && ok {
				trees[root] += trees[factor] * trees[root/factor]
				trees[root] %= MOD
			}
		}
	}
	sum := 0

	for _, v := range trees {
		sum = sum + v
		sum %= MOD
	}
	return sum
}

// func Driver() {
// 	arr := []int{2, 4, 5, 10}
// 	fmt.Println(numFactoredBinaryTrees(arr))
// }
