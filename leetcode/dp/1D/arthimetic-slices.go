// https://leetcode.com/problems/arithmetic-slices/

package dp

func numberOfArithmeticSlices(nums []int) int {
	sum := 0

	dp := make([]int, len(nums))

	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp[i] += dp[i-1] + 1
			sum += dp[i]
		}
	}

	return sum
}

func numberOfArithmeticSlicesMem(nums []int) int {
	sum := 0
	dp := 0

	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp++
			sum += dp
		} else {
			dp = 0
		}
	}

	return sum
}
