package main

import "fmt"

func main() {
	quant := trappedWater([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	fmt.Println(quant)
}

func trappedWater(arr []int) int {
	var left []int
	var right []int

	low := 0
	for _, val := range arr {
		if low < val {
			low = val
		}
		left = append(left, low)
	}
	low = 0
	for i := len(arr) - 1; i >= 0; i-- {
		if low < arr[i] {
			low = arr[i]
		}
		right = append([]int{low}, right...)
	}
	quantity := 0
	for i, val := range arr {
		max := 0
		if left[i] > right[i] {
			max = right[i]
		} else {
			max = left[i]
		}
		quantity += max - val
	}
	return quantity
}
