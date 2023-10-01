package main

import (
	"basic"
	"fmt"
)

func main() {
	// arr := []int{1, 2, 3, 4, 5}
	// fmt.Println(basic.GetSum(arr))

	// arr := []int{3, 2, 1}
	h := basic.Constructor()
	h.Put(4, 3)
	fmt.Println(h)
}
