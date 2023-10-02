package gametheory

// https://leetcode.com/problems/nim-game/

func CanWinNim(n int) bool {
	if n == 0 {
		return false
	}

	if n < 4 {
		return true
	}

	dp := make([]bool, n+1)

	dp[0] = false
	dp[1], dp[2], dp[3] = true, true, true

	for i := 4; i <= n; i++ {
		dp[i] = !(dp[i-1] && dp[i-2] && dp[i-3])
	}

	return dp[n]
}

func CanWinNim2(n int) bool {
	return n%4 != 0
}
