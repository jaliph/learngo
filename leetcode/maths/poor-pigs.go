package maths

import "math"

func PoorPigs_1L(buckets int, minutesToDie int, minutesToTest int) int {
	testRound := minutesToTest / minutesToDie
	testGroup := testRound + 1
	res := int(math.Ceil(math.Log2(float64(buckets)) / math.Log2(float64(testGroup))))
	return res
}

func PoorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	res := 0
	// base needed
	mul := 1
	base := (minutesToTest / minutesToDie) + 1
	for mul < buckets {
		mul *= base
		res++
	}
	return res
}
