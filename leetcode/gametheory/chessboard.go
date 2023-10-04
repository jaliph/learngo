// https://www.hackerrank.com/challenges/a-chessboard-game-1/problem?isFullScreen=true

package gametheory

import "fmt"

var dp map[string]bool

func giveState(x, y int32) bool {
	if x < 1 || y < 1 || x > 15 || y > 15 {
		return true // this position true, thus the border position is false / terminal position
	}

	if v, ok := dp[fmt.Sprintf("%d#%d", x, y)]; ok {
		return v
	}

	ans := true
	ans = ans && giveState(x-2, y+1)
	ans = ans && giveState(x-2, y-1)
	ans = ans && giveState(x+1, y-2)
	ans = ans && giveState(x-1, y-2)
	ans = !ans
	dp[fmt.Sprintf("%d#%d", x, y)] = ans
	return ans
}

func ChessboardGame(x int32, y int32) string {
	dp = make(map[string]bool)
	if giveState(x, y) {
		return "First"
	} else {
		return "Second"
	}
}
