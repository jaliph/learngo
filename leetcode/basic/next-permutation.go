package basic

func NextPermutation(nums []int) {
	// find the break point arr[i] < arr[i + 1] from the right
	// 1 2 3 5 > 1 2 5 3
	// if break point doesnt exist, reverse the full number
	// find the largest in the right of the breakpoint
	// swap break point with the max
	// reverse the right of the array

	reverse := func(arr *[]int, l, r int) {
		a := *arr
		for l <= r {
			a[l], a[r] = a[r], a[l]
			l++
			r--
		}
	}

	idx := -1
	n := len(nums)

	for i := n - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			idx = i
			break
		}
	}
	if idx == -1 {
		reverse(&nums, 0, n-1)
		return
	}

	idx2 := -1
	for i := n - 1; i >= idx; i-- {
		if nums[i] > nums[idx] {
			idx2 = i
			break
		}
	}

	nums[idx], nums[idx2] = nums[idx2], nums[idx]
	reverse(&nums, idx+1, n-1)
}
