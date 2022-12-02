package main

import "fmt"

func main() {
	num := longestBand([]int{1, 9, 3, 0, 18, 5, 2, 4, 10, 7, 12, 6})
	fmt.Println(num)
}

func longestBand(arr []int) int {
	var longest int = 1

	unorderedSet := make(map[int]bool)

	for _, val := range arr {
		unorderedSet[val] = true
	}

	for i := 0; i < len(arr); i++ {
		parent := arr[i] - 1
		if _, ok := unorderedSet[parent]; !ok {
			cnt := 1
			j := arr[i] + 1
			for ok := unorderedSet[j]; ok; {
				j++
				cnt++
				ok = unorderedSet[j]
			}
			if cnt > longest {
				longest = cnt
			}
		}
	}

	return longest
}
