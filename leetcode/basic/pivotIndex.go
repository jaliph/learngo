package basic

import "query"

func PivotIndex(nums []int) int {
	preFixSum := make([]int, len(nums))

	preFixSum[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		preFixSum[i] = preFixSum[i-1] + nums[i]
	}

	sum := func(l, r int) int {
		return preFixSum[r] - preFixSum[l-1]
	}

	for i := 0; i < len(nums); i++ {
		if i == 0 && sum(i+1, len(nums)-1) == 0 {
			return i
		}

		if i == len(nums)-1 && sum(0, i-1) == 0 {
			return i
		}

		if sum(0, i-1) == sum(i+1, len(nums)-1) {
			return i
		}
	}
	return -1
}

func PivotIndexFenwick(nums []int) int {
	f := query.NewFenwick(len(nums))
	for i, v := range nums {
		f.Add(i, v)
	}

	for i := 0; i < len(nums); i++ {
		if i == 0 && f.Sum(i+1, len(nums)-1) == 0 {
			return i
		}

		if i == len(nums)-1 && f.Sum(0, i-1) == 0 {
			return i
		}

		if f.Sum(0, i-1) == f.Sum(i+1, len(nums)-1) {
			return i
		}
	}
	return -1
}
