package maths

import (
	"strings"
)

// https://leetcode.com/problems/integer-to-roman/
func IntToRoman(num int) string {
	rMap := make(map[int]string)

	order := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	rMap[1] = "I"
	rMap[4] = "IV"
	rMap[5] = "V"
	rMap[9] = "IX"
	rMap[10] = "X"
	rMap[50] = "L"
	rMap[40] = "XL"
	rMap[90] = "XC"
	rMap[100] = "C"
	rMap[400] = "CD"
	rMap[500] = "D"
	rMap[900] = "CM"
	rMap[1000] = "M"

	str := ""
	i := len(order) - 1
	for num > 0 {
		if num/order[i] > 0 {
			str += strings.Repeat(rMap[order[i]], num/order[i])
		}
		num %= order[i]
		i--
	}

	return str
}
