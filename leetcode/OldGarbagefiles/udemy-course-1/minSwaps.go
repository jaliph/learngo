package main

import (
	"fmt"
	"sort"
)

func main() {
	swaps := minSwaps([]int{5, 4, 3, 2, 1})
	fmt.Println(swaps)
	swaps = minSwaps([]int{2, 4, 5, 1, 3})
	fmt.Println(swaps)
}

func minSwaps(arr []int) int {
	newArr := make([]int, len(arr))
	visited := make([]bool, len(arr))
	copy(newArr, arr)
	sort.Ints(newArr)

	fmt.Println(newArr)
	fmt.Println(arr)
	swaps := 0
	for i := 0; i < len(arr); i++ {
		// if position matches or already visited
		if arr[i] == newArr[i] || visited[i] {
			continue
		} else {

			node := i
			cycle := 0
			for !visited[node] {
				cycle++
				visited[node] = true
				node = indexOf(arr[node], newArr)
				fmt.Println("Current ", i, "Next ", node)
				// fmt.Print(i, "->", node)
			}
			fmt.Println("Cycle ", cycle-1)
			swaps += (cycle - 1)
		}
	}
	return swaps
}

func indexOf(element int, data []int) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}
