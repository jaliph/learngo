package dp

// https://leetcode.com/problems/arithmetic-slices-ii-subsequence/description/

func numberOfArithmeticSlices2(nums []int) int {
	total := 0
	dp := make([]map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		dp[i] = map[int]int{}
		for j := 0; j < i; j++ {
			diff := nums[i] - nums[j]

			dp[i][diff]++

			if _, ok := dp[j][diff]; ok {
				dp[i][diff] += dp[j][diff]
				total += dp[j][diff]
			}
		}
	}

	return total
}

// func Driver() {
// 	nums := []int{2, 4, 6, 8, 10}
// 	fmt.Println(numberOfArithmeticSlices2(nums))
// }
