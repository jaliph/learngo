package dp

func MaxProfit(prices []int) int {
	var profit int
	minPrice := int(^uint(0) >> 1)

	for _, p := range prices {
		minPrice = min(minPrice, p)
		profit = max(profit, p-minPrice)
	}

	return profit
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
