package recursion

func Combinations(nums []int) [][]int {
	res := [][]int{}

	var comb func(idx int, stack *[]int)
	comb = func(idx int, stack *[]int) {
		// base
		if idx == len(nums) {
			c := make([]int, len(*stack))
			copy(c, *stack)
			res = append(res, c)
			return
		}

		*stack = append(*stack, nums[idx])
		comb(idx+1, stack)
		*stack = (*stack)[:len(*stack)-1]
		comb(idx+1, stack)
	}

	comb(0, &[]int{})

	return res
}
