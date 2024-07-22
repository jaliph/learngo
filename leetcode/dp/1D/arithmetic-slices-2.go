package dp

// https://leetcode.com/problems/arithmetic-slices-ii-subsequence/description/

func NumberOfArithmeticSlicesIntuitive(nums []int) int {

	res := 0
	n := len(nums)
	dp := make([]map[int]int, len(nums))

	for i := 0; i < n; i++ {
		dp[i] = map[int]int{}
		for j := 0; j < i; j++ {
			diff := nums[i] - nums[j]
			dp[i][diff] += 1 + dp[j][diff]
			res += 1 + dp[j][diff]
		}
	}
	return res - ((n * (n - 1)) / 2) // remove all pairs of 2 length
}

func NumberOfArithmeticSlices2(nums []int) int {
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
