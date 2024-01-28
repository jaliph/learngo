package dp

// https://leetcode.com/problems/longest-common-subsequence

func longestCommonSubsequence(text1 string, text2 string) int {

	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	prev := make([]int, len(text2)+1)

	for i := 0; i < len(text1); i++ {
		dp := make([]int, len(text2)+1)
		for j := 0; j < len(text2); j++ {
			if text1[i] == text2[j] {
				dp[j+1] = prev[j] + 1
			} else {
				dp[j+1] = Max(prev[j+1], dp[j])
			}
		}
		prev = dp
	}

	return prev[len(text2)]
}
