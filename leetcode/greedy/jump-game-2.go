package greedy

func jump(nums []int) int {
	l, r := 0, 0
	steps := 0
	var farthest int
	for r < len(nums)-1 {
		farthest = 0
		for i := l; i <= r; i++ {
			farthest = Max(farthest, i+nums[i])
		}
		l = r + 1
		r = farthest
		steps++
	}
	return steps
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
