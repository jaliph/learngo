// https://www.hackerrank.com/challenges/game-of-stones-1
package gametheory

func GameOfStones(n int32) string {
	if n%7 <= 1 {
		return "Second"
	} else {
		return "First"
	}
}

func GameOfStones2(n int32) string {
	dp := make([]bool, n+1)
	dp[0] = false

	var i int32
	for i = 1; i <= n; i++ {
		fn := true
		if i-2 >= 0 {
			fn = fn && dp[i-2]
		}
		if i-3 >= 0 {
			fn = fn && dp[i-3]
		}
		if i-5 >= 0 {
			fn = fn && dp[i-5]
		}
		dp[i] = !fn
	}

	if dp[n] {
		return "First"
	} else {
		return "Second"
	}
}
