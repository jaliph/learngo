package main

import (
	"fmt"
	"sort"
)

func main() {
	triplets := findtriplets([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15}, 18)
	fmt.Println(triplets)
}

func findtriplets(s []int, target int) [][]int {
	sort.Ints(s)
	var result [][]int
	for i := 0; i < len(s)-2; i++ {
		j := i + 1
		k := len(s) - 1
		for j < k {
			if target-s[i] == s[j]+s[k] {
				result = append(result, []int{s[i], s[j], s[k]})
				j++
				k--
			} else if target-s[i] < s[j]+s[k] {
				k--
			} else {
				j++
			}
		}
	}
	return result
}
