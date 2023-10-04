// https://www.hackerrank.com/challenges/nimble-game-1/problem?isFullScreen=true&h_r=next-challenge&h_v=zen&h_r=next-challenge&h_v=zen&h_r=next-challenge&h_v=zen

package gametheory

func NimbleGame(s []int32) string {
	// Write your code here
	nim_sum := 0
	for idx, v := range s {
		v = v % 2
		if v > 0 {
			nim_sum ^= idx
		}
	}

	if nim_sum > 0 {
		return "First"
	} else {
		return "Second"
	}
}
