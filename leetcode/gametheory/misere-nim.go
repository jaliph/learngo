// https://www.hackerrank.com/challenges/misere-nim-1/problem?isFullScreen=true&h_r=next-challenge&h_v=zen&h_r=next-challenge&h_v=zen

package gametheory

func MisereNim(s []int32) string {
	// Write your code here
	var nim_sum int32 = 0

	freq := make(map[int32]bool)
	for _, i := range s {
		nim_sum = nim_sum ^ i
		freq[i] = true
	}

	if len(freq) == 1 && freq[int32(1)] {
		if len(s)&1 == 1 { // if all are 1s and the if the first move 1 makes, last move 1 will make so 2 wins
			return "Second"
		} else { // else 1 moves first move, last move 2 will make, thus 1 wins
			return "First"
		}
	}

	// basic nim sum
	if nim_sum > 0 {
		return "First"
	} else {
		return "Second"
	}
}
