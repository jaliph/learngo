package main

import "fmt"

func main() {
	fmt.Println(isSubset("coding minutes", "kines"))
}

func isSubset(s1, s2 string) bool {
	i := len(s1) - 1
	j := len(s2) - 1

	for i >= 0 {
		if j < 0 {
			return true
		}
		if s1[i] == s2[j] {
			i--
			j--
		} else {
			i--
		}
	}
	return j < 0
}
