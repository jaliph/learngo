package dp

// TLE  - Check monotonic increasing stacks
func constrainedSubsetSum(nums []int, k int) int {
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
	n := len(nums)
	var globalMax int = -1e9
	memoise := map[int]int{}
	var solve func(int) int
	solve = func(idx int) int {
		if idx >= n {
			return 0
		}

		if v, ok := memoise[idx]; ok {
			return v
		}

		ans := nums[idx]
		for j := idx + 1; j < Min(idx+k+1, n); j++ {
			ans = Max(ans, nums[idx]+solve(j))
		}
		globalMax = Max(ans, globalMax)
		memoise[idx] = ans
		return ans
	}

	solve(0)
	return globalMax
}

// func Driver() {
// 	// nums := []int{-5266, 4019, 7336, -3681, -5767}
// 	// k := 2

// 	// fmt.Println(constrainedSubsetSum(nums, k))

// 	// fmt.Println(constrainedSubsetSum([]int{10, -2, -10, -5, 20}, 2))
// 	fmt.Println(constrainedSubsetSum([]int{10, 2, -10, 5, 20}, 2))
// }
