package main

import "fmt"

func main() {

	arr := []int{10, 5, 3, 2, -6, 9, 11}
	pairs := pairSum(arr, 4)
	if len(pairs) == 0 {
		fmt.Println("No result found")
	} else {
		fmt.Println(pairs)
	}

}

func pairSum(s []int, target int) []int {
	hash := make(map[int]int)
	for _, val := range s {
		if _, ok := hash[target-val]; ok {
			return []int{target - val, val}
		} else {
			hash[val] = target - val
		}
	}
	return []int{}
}
