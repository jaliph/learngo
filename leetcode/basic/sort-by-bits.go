package basic

import (
	"sort"
)

var countMap map[int]int

func SortByBits(arr []int) []int {
	countMap = map[int]int{}
	sort.Slice(arr, func(i, j int) bool {
		c1 := countBits(arr[i])
		c2 := countBits(arr[j])
		if c1 == c2 {
			return arr[i] < arr[j]
		}
		return c1 < c2
	})
	return arr
}

func countBits(n int) int {
	if v, ok := countMap[n]; ok {
		return v
	}
	count := 0
	for n > 0 {
		n = n & (n - 1)
		count++
	}
	return count
}

// func Driver() {
// 	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
// 	x := sortByBits(arr)
// 	fmt.Println(x)
// }
