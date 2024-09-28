package main

import (
	"fmt"
	"heap"
)

func main() {
	x := []int{4, 5, 8, 2}
	kth := heap.NewKthLargest(3, x)
	fmt.Println(kth)
	fmt.Println(kth.Add(3))
	fmt.Println(kth)
	fmt.Println(kth.Add(5))
	fmt.Println(kth.Add(10))
	fmt.Println(kth.Add(9))
	fmt.Println(kth.Add(4))
}
