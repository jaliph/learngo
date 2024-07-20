package dp

func MaxUncrossedLines(nums1 []int, nums2 []int) int {

	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	prev := make([]int, len(nums1)+1)

	for i := 0; i < len(nums2); i++ {
		dp := make([]int, len(nums1)+1)
		for j := 0; j < len(nums1); j++ {
			if nums2[i] == nums1[j] {
				dp[1+j] = 1 + prev[j]
			} else {
				dp[1+j] = Max(prev[1+j], dp[j])
			}
		}
		prev = dp
	}

	return prev[len(nums1)]
}
