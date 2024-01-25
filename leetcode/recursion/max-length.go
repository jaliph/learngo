package recursion

import "fmt"

// https://leetcode.com/problems/maximum-length-of-a-concatenated-string-with-unique-characters

func maxLength(arr []string) int {
	m := map[rune]int{}

	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	add := func(str string, freqMap map[rune]int) {
		for _, v := range str {
			freqMap[v]++
		}
	}

	remove := func(str string, freqMap map[rune]int) {
		for _, v := range str {
			freqMap[v]--
			if freqMap[v] == 0 {
				delete(freqMap, v)
			}
		}
	}

	isOverlap := func(str string, freqMap map[rune]int) bool {
		if len(freqMap) == 0 {
			return false
		}
		for _, v := range str {
			if _, ok := freqMap[v]; ok {
				return true
			}
		}
		return false
	}

	var solve func(int) int
	solve = func(idx int) int {
		if idx >= len(arr) {
			return len(m)
		}

		res := 0
		if !isOverlap(arr[idx], m) {
			add(arr[idx], m)
			res = solve(idx + 1)
			remove(arr[idx], m)
		}
		res = Max(res, solve(idx+1))
		return res
	}

	return solve(0)
}

func Driver() {
	fmt.Println(maxLength([]string{"aa", "bb"}))
	fmt.Println(maxLength([]string{"un", "iq", "ue"}))
}
