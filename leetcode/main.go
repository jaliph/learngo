package main

import (
	"dp"
	"fmt"
)

func main() {
	fmt.Println(dp.MaxDotProductDP([]int{2, 1, -2, 5}, []int{3, 0, -6}))
	fmt.Println(dp.MaxDotProductBrute([]int{3, -2}, []int{2, -6, 7}))
}
