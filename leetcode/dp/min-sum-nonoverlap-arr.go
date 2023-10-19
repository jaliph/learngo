package dp

// https://leetcode.com/problems/find-two-non-overlapping-sub-arrays-each-with-target-sum/solutions/688689/c-general-dp-solution-n-non-overlapping-sub-arrays/

func minSumOfLengths(arr []int, target int) int {
	c := 3 // no of subarray + 1
	r := len(arr) + 1

	const INF int = 1e9

	dp := make([][]int, r)
	rows := make([]int, r*c)

	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := 0; i < len(rows); i++ {
		rows[i] = INF
	}
	for i := range dp {
		dp[i] = rows[i*c : (i+1)*c]
	}

	for i := 0; i <= len(arr); i++ {
		dp[i][0] = 0
	}

	prefixMap := map[int]int{}
	prefixMap[0] = 0

	currSum := 0
	for i := 1; i <= len(arr); i++ {
		d := -1
		currSum += arr[i-1]
		prefixMap[currSum] = i

		if v, ok := prefixMap[currSum-target]; ok {
			d = v
		}

		for j := 1; j <= 2; j++ {
			dp[i][j] = Min(dp[i][j], dp[i-1][j])
			if d != -1 {
				dp[i][j] = Min(dp[i][j], dp[d][j-1]+i-d)
			}
		}
	}

	return dp[len(arr)][2]

}
