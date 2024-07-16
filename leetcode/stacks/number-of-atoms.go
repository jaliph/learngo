package stacks

import (
	"fmt"
	"sort"
	"strconv"
)

// https://leetcode.com/problems/number-of-atoms/

func countOfAtoms(formula string) string {

	isLower := func(ch byte) bool {
		if ch >= 'a' && ch <= 'z' {
			return true
		}
		return false
	}

	isDigit := func(ch byte) bool {
		if ch >= '0' && ch <= '9' {
			return true
		}
		return false
	}

	stack := []map[string]int{}
	stack = append(stack, map[string]int{})

	i := 0
	n := len(formula)
	for i < n {
		if formula[i] == '(' {
			stack = append(stack, map[string]int{})
			i++
		} else if formula[i] == ')' {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			i++ // skip '('
			start := i
			for i < n && isDigit(formula[i]) {
				i++
			}

			multplier := 1
			if start != i {
				multplier, _ = strconv.Atoi(formula[start:i])
			}

			for str, count := range top {
				stack[len(stack)-1][str] = stack[len(stack)-1][str] + (count * multplier)
			}
		} else {
			start := i
			i++
			for i < n && isLower(formula[i]) {
				i++
			}
			element := formula[start:i]

			count := 1
			start = i
			for i < n && isDigit(formula[i]) {
				i++
			}
			if i != start {
				count, _ = strconv.Atoi(formula[start:i])
			}
			stack[len(stack)-1][element] = stack[len(stack)-1][element] + count
		}
	}

	// fmt.Println(stack)
	res := stack[len(stack)-1]

	keys := []string{}
	for k := range res {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	// fmt.Println(keys)
	out := ""
	for _, v := range keys {
		out += v
		if res[v] > 1 {
			out += fmt.Sprintf("%d", res[v])
		}
	}

	return out
}

func Driver() {
	str := "K4(ON(SO3)2)2"
	fmt.Println(countOfAtoms(str))
}
