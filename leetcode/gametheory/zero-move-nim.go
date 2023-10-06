package gametheory

import "fmt"

// https://www.hackerrank.com/challenges/zero-move-nim
func Grundy_ZeroNim(n int, zeroDone bool) int {
	if n == 0 {
		return 0
	}
	m := make(map[int]bool)
	if zeroDone {
		g := Grundy_ZeroNim(n, !zeroDone)
		m[g] = true
	}

	for i := 0; i < n; i++ {
		g := Grundy_ZeroNim(i, zeroDone)
		m[g] = true
	}

	mex := 0
	for i := 0; ; i++ {
		if _, ok := m[i]; !ok {
			mex = i
			break
		}
	}
	return mex
}

func ZeroMoveNim(p []int32) string {
	// Write your code here

	for i := 0; i < 20; i++ {
		fmt.Println(i, " has ", Grundy_ZeroNim(i, true))
	}
	return ""
}
