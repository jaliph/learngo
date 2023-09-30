package main

import (
	"cyclic"
	"fmt"
)

func main() {
	// arr := []int{1, 2, 3, 4, 5}
	// fmt.Println(basic.GetSum(arr))

	arr := []int{4, 3, 2, 1}
	fmt.Println(cyclic.FindDisappearedNumbers(arr))

}
