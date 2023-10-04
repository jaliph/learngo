// https://leetcode.com/problems/find-missing-observations/submissions/
package maths

func MissingRolls(rolls []int, mean int, n int) []int {
	sum := 0
	for _, v := range rolls {
		sum += v
	}

	rem := (mean * (len(rolls) + n)) - sum

	res := []int{}
	if n*1 > rem || rem > n*6 {
		return res
	}

	// Greedy approach
	for rem > 0 {
		dice := min(rem-n+1, 6)
		res = append(res, dice)
		rem -= dice
		n--
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
