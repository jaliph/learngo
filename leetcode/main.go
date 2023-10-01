package main

import (
	"fmt"
	"heap"
)

func main() {
	// arr := []int{1, 2, 3, 4, 5}
	// fmt.Println(basic.GetSum(arr))

	// arr := []int{3, 2, 1}
	h := heap.NewHeap(nil)

	h.Push(5)
	h.Push(2)
	h.Push(100)
	fmt.Println(h)
	fmt.Println(h.Peek())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h)

}
