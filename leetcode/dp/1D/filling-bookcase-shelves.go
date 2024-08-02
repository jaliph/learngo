package dp

// https://leetcode.com/problems/filling-bookcase-shelves

func MinHeightShelves(books [][]int, shelfWidth int) int {
	const INF int = 1e9 + 7
	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}
	var dfs func(idx int) int

	memoise := make([]int, len(books)+1)
	for i := range memoise {
		memoise[i] = -1
	}
	memoise[len(books)] = 0

	dfs = func(idx int) int {

		// if idx == len(books) {
		// 	return 0
		// }
		if memoise[idx] != -1 {
			return memoise[idx]
		}

		currWidth := 0
		maxHeight := 0
		memoise[idx] = INF

		for j := idx; j < len(books); j++ {
			if currWidth+books[j][0] <= shelfWidth {
				currWidth += books[j][0]
				maxHeight = Max(maxHeight, books[j][1])
				memoise[idx] = Min(memoise[idx], dfs(j+1)+maxHeight)
			} else {
				break
			}
		}
		return memoise[idx]
	}

	return dfs(0)
}
