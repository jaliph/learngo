package dp

import "math"

// https://leetcode.com/problems/palindrome-partitioning-ii/
func MinCut(str string) int {
	n := len(str)
	memo := make(map[int]int)
	return f(0, n, str, memo) - 1
}

func isPalindrome(i, j int, s string) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func f(i, n int, str string, memo map[int]int) int {
	// Base case:
	if i == n {
		return 0
	}

	// Check memo:
	if val, found := memo[i]; found {
		return val
	}

	minCost := math.MaxInt32
	// String[i...j]
	for j := i; j < n; j++ {
		if isPalindrome(i, j, str) {
			cost := 1 + f(j+1, n, str, memo)
			if cost < minCost {
				minCost = cost
			}
		}
	}

	// Store in memo:
	memo[i] = minCost

	return minCost
}
