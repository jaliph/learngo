package stacks

func LargestRectangleArea(heights []int) int {

	stack := []int{}
	res := 0

	Max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	for i := 0; i < len(heights); i++ {
		start := i
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			area := heights[top] * (i - top)
			res = Max(res, area)
			start = top
		}
		stack = append(stack, start)
	}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		area := heights[top] * (len(heights) - top)
		res = Max(res, area)
	}

	return res
}
