package greedy

// https://leetcode.com/problems/minimum-difference-between-largest-and-smallest-value-in-three-moves/

func MinDifference(nums []int) int {
	if len(nums) <= 4 {
		return 0
	}
	const INF int = 1e9 + 7
	Min1, Min2, Min3, Min4 := INF, INF, INF, INF
	Max1, Max2, Max3, Max4 := -INF, -INF, -INF, -INF

	for _, n := range nums {
		if n < Min1 {
			Min4 = Min3
			Min3 = Min2
			Min2 = Min1
			Min1 = n
		} else if n < Min2 {
			Min4 = Min3
			Min3 = Min2
			Min2 = n
		} else if n < Min3 {
			Min4 = Min3
			Min3 = n
		} else if n < Min4 {
			Min4 = n
		}

		if n > Max1 {
			Max4 = Max3
			Max3 = Max2
			Max2 = Max1
			Max1 = n
		} else if n > Max2 {
			Max4 = Max3
			Max3 = Max2
			Max2 = n
		} else if n > Max3 {
			Max4 = Max3
			Max3 = n
		} else if n > Max4 {
			Max4 = n
		}
	}
	// fmt.Println(Min1, Min2, Min3, Min4)
	// fmt.Println(Max1, Max2, Max3, Max4)

	Min := func(nums []int) int {
		// fmt.Println(nums)
		min := nums[0]
		for _, v := range nums {
			if min > v {
				min = v
			}
		}
		return min
	}
	return Min([]int{Max4 - Min1, Max3 - Min2, Max2 - Min3, Max1 - Min4})
}
