// https://leetcode.com/problems/plus-one
package maths

func PlusOne(digits []int) []int {
	carry := 0
	incr := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digitsum := digits[i] + incr + carry
		incr = 0
		carry = digitsum / 10
		digits[i] = digitsum % 10
	}
	if carry > 0 {
		digits = append([]int{carry}, digits...)
	}
	return digits
}
