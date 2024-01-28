package dp

// duplicate to stacks

// https://leetcode.com/problems/sum-of-subarray-minimums/solutions/4595290/c-python-dp-monotonic-stack-23ms-beats-100/?envType=daily-question&envId=2024-01-20

func sumSubarrayMins(arr []int) int {
	const MOD int = 1e9 + 7
	dp := make([]int, len(arr))
	ans := 0

	stack := []int{}
	for i := 0; i < len(arr); i++ {
		for len(stack) > 0 && stack[len(stack)-1] >= arr[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			j := stack[len(stack)-1]
			dp[i] = dp[j] + (arr[i] * (i - j))
		} else {
			dp[i] = arr[i] * (i + 1)
		}
		stack = append(stack, i)
		ans += dp[i]
		ans %= MOD
	}

	return ans
}
