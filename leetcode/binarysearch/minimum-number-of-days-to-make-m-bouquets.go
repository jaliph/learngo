package binarysearch

import "slices"

func countBouquets(bloomDay []int, day int, k int) int {
	cnt := 0
	res := 0

	for _, val := range bloomDay {
		if val <= day {
			cnt++
			if cnt == k {
				res++
				cnt = 0
			}
		} else {
			cnt = 0
		}
	}

	return res
}

func MinDays(bloomDay []int, m int, k int) int {
	l := slices.Min(bloomDay)
	r := slices.Max(bloomDay)

	res := -1

	for l <= r {
		mid := l + (r-l)/2
		if countBouquets(bloomDay, mid, k) >= m {
			res = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return res
}
