package recursion

import "fmt"

// https://leetcode.com/problems/maximum-length-of-a-concatenated-string-with-unique-characters

func maxLength(arr []string) int {
	const all1s int = 1<<27 - 1
	charArray := make([]int, len(arr))
	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i, v := range arr {
		for _, ch := range v {
			if charArray[i]&(1<<(ch-'a')) > 0 {
				charArray[i] = all1s
			}
			charArray[i] = charArray[i] | 1<<(ch-'a')
		}
	}

	get1Counts := func(i int) int {
		count := 0
		for i > 0 {
			i = i & (i - 1)
			count++
		}
		return count
	}

	var recur func(int, int) int
	recur = func(idx int, bit int) int {
		if idx >= len(charArray) {
			return get1Counts(bit)
		}

		if charArray[idx] == all1s {
			return recur(idx+1, bit)
		}

		res := 0
		if bit&charArray[idx] == 0 {

			res = recur(idx+1, bit|charArray[idx])
		}
		return Max(res, recur(idx+1, bit))
	}

	return recur(0, 0)
}

func Driver() {
	// fmt.Println(maxLength([]string{"aa", "bb"}))
	fmt.Println(maxLength([]string{"a", "b", "c"}))
}
