package greedy

import (
	"sort"
)

// https://leetcode.com/problems/eliminate-maximum-number-of-monsters/description/?envType=daily-question&envId=2023-11-07

func eliminateMaximum(dist []int, speed []int) int {
	time := []int{}
	res := 0

	Ceil := func(a, b int) int {
		return (a + b - 1) / b
	}

	for i := range dist {
		time = append(time, Ceil(dist[i], speed[i]))
	}

	sort.Slice(time, func(i, j int) bool {
		return time[i] < time[j]
	})

	for i := 0; i < len(time); i++ {
		if time[i] > i {
			res++
		} else {
			return res
		}
	}

	return res
}

// func Driver() {
// 	dist := []int{1, 3, 4}
// 	speed := []int{1, 1, 1}
// 	fmt.Println(eliminateMaximum(dist, speed))
// }
