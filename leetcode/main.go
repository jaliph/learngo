package main

import (
	"basic"
	"fmt"
)

func main() {
	arr := []int{1, 1, 1, 4}
	fmt.Println(arr, basic.FindDuplicate(arr))
}
