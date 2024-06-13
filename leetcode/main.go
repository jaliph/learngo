package main

import (
	"dp/dp-strings"
	"fmt"
)

func main() {
	// nc := democracy.NewConfig(
	// 	democracy.WithMaxPacketSize(50),
	// 	democracy.WithPeers([]string{"127.0.0.1:12345"}),
	// )
	// _, ch := democracy.NewDemocracy(nc)
	// <-ch
	fmt.Println(dp.IsInterleave_TD("aabcc", "dbbca", "aadbbcbcac"))

}

// do it the other way
