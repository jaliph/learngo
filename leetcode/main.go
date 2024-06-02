package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(customSortString("bcafg", "abcd"))
}

func customSortString(order string, s string) string {
	oMap := map[byte]int{}
	for k, v := range []byte(order) {
		oMap[v] = k + 1
	}

	str := []byte(s)
	sort.Slice(str, func(i, j int) bool {
		oi := len(order) + 1
		oj := len(order) + 1
		if v, ok := oMap[str[i]]; ok {
			oi = v
		}
		if v, ok := oMap[str[j]]; ok {
			oj = v
		}
		return oi < oj
	})

	return string(str[:])
}

// do it the other way
