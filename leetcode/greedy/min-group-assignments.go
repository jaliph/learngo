package greedy

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/minimum-number-of-groups-to-create-a-valid-assignment/

func minGroupsForValidAssignment(nums []int) int {
	count := make(map[int]int)

	for _, num := range nums {
		count[num]++
	}

	const INF int = 1e9
	minFreq := INF
	listFreq := []int{}

	for _, v := range count {
		if minFreq > v {
			minFreq = v
		}
		listFreq = append(listFreq, v)
	}

	if len(listFreq) == 1 {
		return 1
	}
	res := len(nums)

	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	canPartition := func(guessedParition int, frequencies []int) int {
		result := 0

		for _, freq := range frequencies {
			partitionCount := freq / guessedParition
			remainder := freq % guessedParition

			if remainder > partitionCount {
				return INF
			}

			result += int(math.Ceil(float64(freq) / float64(guessedParition+1)))
		}
		// fmt.Println(guessedParition, result)
		return result
	}

	for i := 1; i <= minFreq; i++ {
		res = Min(res, canPartition(i, listFreq))
	}

	return res
}

func Driver() {
	fmt.Println(minGroupsForValidAssignment([]int{10, 10, 10, 3, 1, 1}))
}
