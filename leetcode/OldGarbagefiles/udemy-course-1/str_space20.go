package main

import (
	"fmt"
)

func main() {
	str := "hello are you ?"
	fmt.Println(string(insert20([]rune(str))))
}

func insert20(str []rune) []rune {
	spaces := 0
	for _, v := range str {
		if string(v) == " " {
			spaces++
		}
	}

	idx := (spaces * 2) + len(str) - 1
	orig := len(str) - 1
	newArr := make([]rune, idx+1)
	for idx >= 0 {
		if string(str[orig]) != " " {
			newArr[idx] = str[orig]
			idx--
			orig--
		} else {
			newArr[idx] = rune('0')
			newArr[idx-1] = rune('2')
			newArr[idx-2] = rune('x')
			orig--
			idx = idx - 3
		}
	}
	return newArr
}
