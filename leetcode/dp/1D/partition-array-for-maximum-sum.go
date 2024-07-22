package dp

// https://leetcode.com/problems/partition-array-for-maximum-sum/
func maxSumAfterPartitioning(arr []int, k int) int {
	Max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	Min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}

	n := len(arr)
	memoise := make([]int, n)
	var recur func(int) int
	recur = func(idx int) int {
		if idx >= n {
			return 0
		}

		if memoise[idx] != 0 {
			return memoise[idx]
		}

		max := 0
		ans := 0

		for j := idx; j < Min(n, idx+k); j++ {
			max = Max(max, arr[j])
			ans = Max(ans, max*(j-idx+1)+recur(j+1))
		}
		memoise[idx] = ans
		return ans
	}

	return recur(0)
}
