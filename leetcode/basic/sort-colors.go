package basic

func SortColors(nums []int) {
	var l, r, m int
	l = 0
	m = 0
	r = len(nums) - 1

	swap := func(arr *[]int, i, j int) {
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	}

	for m <= r {
		if nums[m] > 1 {
			swap(&nums, m, r)
			r--
		} else if nums[m] < 1 {
			swap(&nums, l, m)
			l++
			m++
		} else {
			m++
		}
	}
}
