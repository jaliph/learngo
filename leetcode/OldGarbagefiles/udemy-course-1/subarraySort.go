package main

import (
	"fmt"
	"math"
)

func main() {
	start, end := subArraySort([]int{1, 2, 3, 4, 5, 8, 6, 7, 9, 10, 11})
	fmt.Println(start, end)
}

func subArraySort(arr []int) (int, int) {
	var largest int = math.MinInt
	var smallest int = math.MaxInt
	for i := 0; i < len(arr); i++ {
		if outOfOrder(arr, i) {
			smallest = min(smallest, arr[i])
			largest = max(largest, arr[i])
		}
	}

	if smallest == math.MaxInt {
		return -1, -1
	}
	i := 0
	for arr[i] <= smallest {
		i++
	}

	j := len(arr) - 1
	for arr[j] >= largest {
		j--
	}

	return i, j
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func outOfOrder(arr []int, i int) bool {
	if i == 0 {
		return arr[0] > arr[1]
	}
	if i == len(arr)-1 {
		return arr[i] < arr[len(arr)-1]
	}
	return arr[i] > arr[i+1] || arr[i] < arr[i-1]
}
