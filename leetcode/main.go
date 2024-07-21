package main

import (
	"basic"
	"fmt"
)

func main() {
	arr := []int{1, 3, 2}

	basic.NextPermutation(arr)
	fmt.Println(arr)
}
