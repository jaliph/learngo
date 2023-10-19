package dp

import (
	"fmt"
	"strconv"
)

// https://leetcode.com/problems/numbers-with-repeated-digits/

func numDupDigitsAtMostN(n int) int {
	str := ""
	temp := n
	for temp > 0 {
		str = fmt.Sprintf("%d", temp%10) + str
		temp = temp / 10
	}

	convertBool := func(b bool) int {
		if b {
			return 1
		}
		return 0
	}

	dp := map[[3]int]int{}

	var solve func(int, int, int) int
	solve = func(idx int, last int, mask int) int {
		if idx == len(str) {
			// this means none of the digits where chosen in the mask
			if mask == 0 {
				return 0
			} else {
				return 1
			}
		}

		if v, ok := dp[[3]int{idx, last, mask}]; ok {
			return v
		}

		var res int
		var till int
		if last == 1 {
			till, _ = strconv.Atoi(string(str[idx]))
		} else {
			till = 9
		}

		for i := 0; i <= till; i++ {
			// all 0 s in the front, skip this
			if mask == 0 && i == 0 {
				res += solve(idx+1, last&convertBool(i == till), mask)
				continue
			}
			// if not set in mask
			if ((mask >> i) & 1) == 0 {
				res += solve(idx+1, last&convertBool(i == till), (mask | 1<<i))
			}
		}
		dp[[3]int{idx, last, mask}] = res
		return res
	}

	return n - solve(0, 1, 0)
}
