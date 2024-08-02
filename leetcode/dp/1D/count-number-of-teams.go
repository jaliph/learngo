package dp

// https://leetcode.com/problems/count-number-of-teams/

func NumTeams_TD(rating []int) int {
	var dfs func(int, int, bool) int
	dfs = func(i int, count int, ascend bool) int {
		if count == 3 {
			return 1
		}
		if i == len(rating) {
			return 0
		}

		res := 0
		for j := i + 1; j < len(rating); j++ {
			if ascend && rating[j] > rating[i] {
				res += dfs(j, count+1, ascend)
			}
			if !ascend && rating[j] < rating[i] {
				res += dfs(j, count+1, ascend)
			}
		}
		return res
	}
	res := 0
	for i := range rating {
		res += dfs(i, 1, true)
		res += dfs(i, 1, false)
	}
	return res
}

// Greedy
func NumTeamsGR(rating []int) int {

	res := 0
	var leftSmall, rightLarge, leftLarge, rightSmall int
	for i := 1; i < len(rating)-1; i++ {
		leftSmall = 0
		rightLarge = 0

		for j := i - 1; j >= 0; j-- {
			if rating[j] < rating[i] {
				leftSmall++
			}
		}

		for j := i + 1; j < len(rating); j++ {
			if rating[j] > rating[i] {
				rightLarge++
			}
		}

		res += leftSmall * rightLarge

		leftLarge = i - leftSmall
		rightSmall = len(rating) - i - 1 - rightLarge
		res += leftLarge * rightSmall
	}
	return res
}

// 0, 1, 2, 3, 4
