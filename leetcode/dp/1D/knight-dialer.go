package dp

// https://leetcode.com/problems/knight-dialer/

func knightDialer(n int) int {

	dialerMap := map[int][]int{
		0: {4, 6},
		1: {6, 8},
		2: {7, 9},
		3: {4, 8},
		4: {0, 3, 9},
		5: {},
		6: {0, 7, 1},
		7: {2, 6},
		8: {1, 3},
		9: {2, 4},
	}

	dp := make([]int, 10)
	for i := range dp {
		dp[i] = 1
	}
	const MOD = 1e9 + 7

	for i := 2; i <= n; i++ {
		temp := make([]int, 10)
		for j := 0; j <= 9; j++ {
			next := dialerMap[j]
			for _, v := range next {
				temp[j] = temp[j] + dp[v]
				temp[j] %= MOD
			}
		}
		dp = temp
	}

	sum := 0
	for _, v := range dp {
		sum += v
		sum %= MOD
	}
	return sum
}
