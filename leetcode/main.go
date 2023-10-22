package main

import (
	"fmt"
)

// func main() {
// 	dp.LCS()
// }

func minGroupsForValidAssignment(nums []int) int {
	freq := make(map[int]int)
	var maxFreq int = -1e9
	var minFreq int = 1e9

	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for _, num := range nums {
		freq[num]++
		if freq[num] > maxFreq {
			maxFreq = freq[num]
		}
	}
	for _, v := range freq {
		minFreq = Min(minFreq, v)
	}

	fmt.Println(maxFreq, minFreq)
	if maxFreq-minFreq <= 1 {
		return len(freq)
	} else {
		return maxFreq - minFreq
	}
}

func main() {
	nums := []int{10, 10, 10, 3, 1, 1}
	fmt.Println(minGroupsForValidAssignment(nums))
}
