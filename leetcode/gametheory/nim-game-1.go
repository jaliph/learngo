// https://www.hackerrank.com/challenges/nim-game-1/problem?isFullScreen=true&h_r=next-challenge&h_v=zen

package gametheory

func NimGame(pile []int32) string {
	// Write your code here
	var x int32 = 0
	for _, i := range pile {
		x = x ^ i
	}

	if x > 0 {
		return "First"
	} else {
		return "Second"
	}
}
