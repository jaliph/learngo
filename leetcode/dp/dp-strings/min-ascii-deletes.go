package dp

import (
	"fmt"
)

// https://leetcode.com/problems/minimum-ascii-delete-sum-for-two-strings

func minimumDeleteSumTD(s1 string, s2 string) int {
	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	memoise := map[[2]int]int{}

	var solve func(i, j int) int
	solve = func(i, j int) int {

		// base
		if i >= len(s1) && j >= len(s2) {
			return 0
		}

		if v, ok := memoise[[2]int{i, j}]; ok {
			return v
		}

		if j >= len(s2) {
			cost := 0
			for i < len(s1) {
				cost += int(s1[i])
				i++
			}
			memoise[[2]int{i, j}] = cost
			return cost
		}

		if i >= len(s1) {
			cost := 0
			for j < len(s2) {
				cost += int(s2[j])
				j++
			}
			memoise[[2]int{i, j}] = cost
			return cost
		}

		// both are same.. no cost required
		if s1[i] == s2[j] {
			memoise[[2]int{i, j}] = solve(i+1, j+1)
			return memoise[[2]int{i, j}]
		}

		memoise[[2]int{i, j}] = Min(int(s1[i])+solve(i+1, j), int(s2[j])+solve(i, j+1))
		return memoise[[2]int{i, j}]
	}

	return solve(0, 0)
}

func minimumDeleteSum(s1 string, s2 string) int {

	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	prev := make([]int, len(s2)+1)
	for i, v := range s2 {
		prev[i+1] = int(v) + prev[i]
	}
	fmt.Println(prev)

	for i := 0; i < len(s1); i++ {
		dp := make([]int, len(s2)+1)
		dp[0] = int(s1[i]) + prev[0]
		for j := 0; j < len(s2); j++ {
			if s1[i] == s2[j] {
				dp[j+1] = prev[j]
			} else {
				dp[j+1] = Min(int(s1[i])+prev[j+1], int(s2[j])+dp[j])
			}
		}
		// fmt.Println(dp)
		prev = dp
	}

	return prev[len(s2)]
}

func Driver() {
	s1 := "delete"
	s2 := "leet"

	fmt.Println(minimumDeleteSum(s1, s2))
}
