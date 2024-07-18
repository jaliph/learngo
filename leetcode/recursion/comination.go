package recursion

// https://leetcode.com/problems/combinations/

func Combine(n int, k int) [][]int {
	res := [][]int{}

	var dfs func(int, *[]int)
	dfs = func(idx int, arr *[]int) {
		if idx > n+1 {
			return
		}
		if len(*arr) == k {
			combo := make([]int, k)
			copy(combo, *arr)
			res = append(res, combo)
			return
		}

		dfs(idx+1, arr)
		*arr = append(*arr, idx)
		dfs(idx+1, arr)
		*arr = (*arr)[:len(*arr)-1]
	}

	dfs(1, &[]int{})

	return res
}
