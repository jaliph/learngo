//https://leetcode.com/problems/maximum-number-of-events-that-can-be-attended-ii/description/

package dp

import (
	"fmt"
	"sort"
)

func maxValue(events [][]int, k int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})

	prev := [][]int{{0, 0}}
	curr := [][]int{{0, 0}}

	findIndex := func(e [][]int, target int) int {
		l, r := 0, len(e)
		var mid int
		for l < r {
			mid = l + (r-l)/2
			if e[mid][0] < target {
				l = mid + 1
			} else {
				r = mid
			}
		}
		return l
	}

	for K := 0; K < k; K++ {
		for _, v := range events {
			idx := findIndex(prev, v[0]) - 1
			if prev[idx][1]+v[2] > curr[len(curr)-1][1] {
				curr = append(curr, []int{v[1], prev[idx][1] + v[2]})
			}
		}
		fmt.Println(curr)
		prev, curr = curr, [][]int{{0, 0}}
	}

	return prev[len(prev)-1][1]
}

// TLE
func maxValue_TD(events [][]int, k int) int {

	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})

	dp := map[[2]int]int{}
	n := len(events)

	Max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var solve func(int, int, int) int
	solve = func(idx, k, end int) int {
		if idx == n || k == 0 {
			return 0
		}

		if events[idx][0] <= end {
			return solve(idx+1, k, end)
		}

		if v, ok := dp[[2]int{idx, k}]; ok {
			return v
		}

		res := Max(
			events[idx][2]+solve(idx+1, k-1, events[idx][1]),
			solve(idx+1, k, end))

		dp[[2]int{idx, k}] = res
		return res
	}

	return solve(0, k, 0)
}
