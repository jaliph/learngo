package main

import (
	"fmt"
)

func main() {
	height := findHeighestMountain([]int{5, 6, 1, 2, 3, 4, 5, 4, 3, 2, 0, 1, 2, 3, -2, 4})
	fmt.Println(height)
}

func findHeighestMountain(s []int) int {
	var largest int = 0
	for i := 1; i < len(s)-1; {
		if s[i] > s[i-1] && s[i] > s[i+1] {

			cnt := 1
			j := i
			for j >= 1 && s[j] > s[j-1] {
				cnt++
				j--
			}
			for i < len(s)-1 && s[i] > s[i+1] {
				cnt++
				i++
			}
			if largest < cnt {
				largest = cnt
			}
		} else {
			i++
		}
	}
	return largest
}
