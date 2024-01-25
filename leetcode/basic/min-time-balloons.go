package basic

// https://leetcode.com/problems/minimum-time-to-make-rope-colorful/
func MinCost(colors string, neededTime []int) int {

	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}
	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	minTime := 0

	for i := 0; i < len(colors); i++ {
		prev := neededTime[i]
		for i+1 < len(colors) && colors[i] == colors[i+1] {
			minTime += Min(prev, neededTime[i+1])
			prev = Max(prev, neededTime[i+1])
			i++
		}
	}

	return minTime
}
