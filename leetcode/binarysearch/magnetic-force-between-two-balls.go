package binarysearch

import "sort"

// https://leetcode.com/problems/magnetic-force-between-two-balls/

// Agressive Cows from SPOJ
func MaxDistance(position []int, m int) int {
	l := 0
	r := int(1e9)

	sort.Slice(position, func(i, j int) bool {
		return position[i] < position[j]
	})

	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var res int
	for l <= r {
		mid := l + (r-l)/2
		if canPlace(position, mid, m) {
			res = Max(res, mid)
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return res
}

func canPlace(position []int, dist int, balls int) bool {

	count := 1
	initalPosition := position[0]
	for i := 1; i < len(position); i++ {
		if position[i]-initalPosition >= dist {
			count++
			initalPosition = position[i]
		}
	}
	return count >= balls
}
