package main

import (
	"fmt"
)

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	minProd := nums[0]
	maxProd := nums[0]
	result := nums[0]

	for i := 1; i < len(nums); i++ {
		// If current element is zero, we reset minProd and maxProd to 1
		// and consider the current zero as a potential max product (0)
		if nums[i] == 0 {
			minProd = 1
			maxProd = 1
			result = MaxIntV(result, 0)
			continue
		}

		// Store the current maxProd before updating it
		tempMax := maxProd

		// Update maxProd and minProd
		maxProd = MaxIntV(nums[i], maxProd*nums[i], minProd*nums[i])
		minProd = MinIntV(nums[i], tempMax*nums[i], minProd*nums[i])

		// Update result to be the maximum of the current result and the updated maxProd
		result = MaxIntV(result, maxProd)
	}

	return result
}

func MinIntV(args ...int) int {
	min := args[0]
	for _, v := range args {
		if min > v {
			min = v
		}
	}
	return min
}

func MaxIntV(args ...int) int {
	max := args[0]
	for _, v := range args {
		if max < v {
			max = v
		}
	}
	return max
}

func main() {
	nums := []int{0, 10, 10, 10, 10, 10, 10, 10, 10, 10, -10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 0}
	fmt.Println(maxProduct(nums)) // Expected output: 10000000000
}
