// https://leetcode.com/problems/trapping-rain-water/

package twopointer

func trap(height []int) int {
	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	l, r := 0, len(height)-1

	leftMax := height[l]
	rightMax := height[r]
	res := 0

	for l < r {
		if leftMax < rightMax {
			l++
			leftMax = Max(leftMax, height[l])
			res += leftMax - height[l]
		} else {
			r--
			rightMax = Max(rightMax, height[r])
			res += rightMax - height[r]
		}
	}
	return res
}
