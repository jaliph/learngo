package maths

import (
	"fmt"
	"strconv"
)

// https://leetcode.com/problems/multiply-strings/
func MultiplyString(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	n1 := reverseInts(num1)
	n2 := reverseInts(num2)
	res := make([]int, len(n1)+len(n2))

	for i := range n1 {
		for j := range n2 {
			res[i+j] += n1[i] * n2[j]
			res[i+j+1] += res[i+j] / 10
			res[i+j] = res[i+j] % 10
		}
	}

	fmt.Println(res)

	res = reverse(res)
	beg := 0
	for {
		if res[beg] != 0 {
			break
		}
		beg++
	}
	res = res[beg:]
	str := ""
	for i := 0; i < len(res); i++ {
		str += fmt.Sprintf("%d", res[i])
	}
	return str
}

func reverseInts(s string) []int {
	res := make([]int, len(s))
	i := len(s)
	for _, v := range s {
		i--
		res[i], _ = strconv.Atoi(string(v))
	}
	return res
}

func reverse(n []int) []int {
	l, r := 0, len(n)-1
	for l < r {
		n[l], n[r] = n[r], n[l]
		l++
		r--
	}
	return n
}
