package basic

import "fmt"

func countingSort(arr []int) []int {

	MaxV := func(args ...int) int {
		max := args[0]

		for i := 1; i < len(args); i++ {
			if max < args[i] {
				max = args[i]
			}
		}
		return max
	}

	max := MaxV(arr...)
	// countArray :=

	// for _, v := range arr {
	// 	freqMap[v]++
	// }

	fmt.Println(max)
	return []int{}
}

func Driver() {
	countingSort([]int{4, 3, 12, 1, 5, 5, 3, 9})
}
