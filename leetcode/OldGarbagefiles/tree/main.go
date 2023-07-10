package main

import (
	"BT"
	"fmt"
)

func main() {
	v := &BT.BT{}
	x := BT.BTNode{
		Val: 1,
		Left: &BT.BTNode{
			Val: 2,
		},
		Right: &BT.BTNode{
			Val: 3,
		},
	}
	fmt.Println(v)
	fmt.Println(x)
	fmt.Println(v.PreOrderTraversal(&x))

}
