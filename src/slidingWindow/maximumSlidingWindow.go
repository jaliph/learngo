package main

import (
	"fmt"
	"strings"
)

func findMaxSlidingWindow(nums []int, windowSize int) []int {
	result := make([]int, 0)

	if len(nums) == 0 {
		return result
	}

	window := NewDeque()

	if windowSize > len(nums) {
		windowSize = len(nums)
	}

	for i := 0; i < windowSize; i++ {
		if !window.IsEmpty() && nums[i] >= window.Back() {
			window.PopBack()
		}
		window.PushBack(i)
	}
	result = append(result, nums[window.Back()])

	for i := windowSize; i < len(nums); i++ {
		if !window.IsEmpty() && nums[i] >= window.Back() {
			window.PopBack()
		}

		if !window.IsEmpty() && window.Front() <= (i-windowSize) {
			window.PopFront()
		}

		window.PushBack(i)
		result = append(result, nums[window.Front()])
	}

	return result
}

func main() {
	targetList := []int{18, 2, 1, 2, 4, 3, 4, 2}
	numsList := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{10, 6, 9, -3, 23, -1, 34, 56, 67, -1, -4, -8, -2, 9, 10, 34, 67},
		{4, 5, 6, 1, 2, 3},
		{9, 5, 3, 1, 6, 3},
		{2, 4, 6, 8, 10, 12, 14, 16},
		{-1, -1, -2, -4, -6, -7},
		{},
		{4, 4, 4, 4, 4, 4},
	}

	for index, value := range numsList {
		fmt.Printf("%d. Original list:\t%s\n", index+1, strings.Replace(fmt.Sprint(value), " ", ", ", -1))
		fmt.Printf("   Input Window size:\t\t%d\n", targetList[index])
		fmt.Printf("   Window size used for the problem:\t\t%d\n", findMaxSlidingWindow(value, targetList[index]))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
