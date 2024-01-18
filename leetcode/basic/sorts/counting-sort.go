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

	countArr := make([]int, max+1)
	output := make([]int, len(arr))

	// increase freq
	for _, v := range arr {
		countArr[v]++
	}

	fmt.Println(countArr)

	// prefix sum counting array
	for i := 1; i < len(countArr); i++ {
		countArr[i] += countArr[i-1]
	}
	fmt.Println(countArr)

	// set the numbers in the original array from the counted Freq
	for i := len(arr) - 1; i >= 0; i-- {
		output[countArr[arr[i]]-1] = arr[i]
		countArr[arr[i]]--
	}
	fmt.Println(output)
	return output
}

func Driver() {
	countingSort([]int{4, 3, 12, 1, 5, 5, 3, 9})
}
