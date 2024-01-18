// https://leetcode.com/problems/buy-two-chocolates/description/

package basic

func buyChoco(prices []int, money int) int {
	var min1 int = 1e12
	var min2 int = 1e12

	for _, p := range prices {
		if p < min1 {
			min2 = min1
			min1 = p
		} else if min2 > p {
			min2 = p
		}
	}
	if (min1 + min2) <= money {
		return money - (min1 + min2)
	}
	return money
}
