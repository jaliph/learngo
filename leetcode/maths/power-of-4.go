package maths

// https://leetcode.com/problems/power-of-four/

import "math"

func IsPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}
	logbase4 := math.Log(float64(n)) / math.Log(4)

	return math.Mod(logbase4, 1.0) == 0

	if n == 1 {
		return true
	}
	if n <= 0 || n%4 > 0 {
		return false
	}
	return IsPowerOfFour(n / 4)
}
