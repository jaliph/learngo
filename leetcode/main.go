package main

import (
	"fmt"
	"greedy"
)

func main() {
	// nc := democracy.NewConfig(
	// 	democracy.WithMaxPacketSize(50),
	// 	democracy.WithPeers([]string{"127.0.0.1:12345"}),
	// )
	// _, ch := democracy.NewDemocracy(nc)
	// <-ch
	fmt.Println(greedy.MinIncrementForUnique([]int{3, 2, 1, 2, 1, 7}))

}

// do it the other way
