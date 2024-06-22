package main

import (
	"fmt"
	"slidingwindow"
)

func main() {
	count := slidingwindow.NumberOfSubarrays([]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 2)
	fmt.Println(count)
}

// do it the other way
