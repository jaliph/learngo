// https://www.hackerrank.com/challenges/poker-nim-1/problem?isFullScreen=true
package gametheory

func PokerNim(k int32, c []int32) string {
	// Write your code here
	var nim_sum int32 = 0
	for _, val := range c {
		nim_sum ^= val
	}

	if nim_sum > 0 {
		return "First"
	} else {
		return "Second"
	}
}
