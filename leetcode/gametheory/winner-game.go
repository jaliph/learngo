// https://leetcode.com/problems/remove-colored-pieces-if-both-neighbors-are-the-same-color
package gametheory

func WinnerOfGame(colors string) bool {
	aliceScore := 0
	bobScore := 0
	for i := 1; i < len(colors)-1; i++ {
		current := colors[i]
		prev := colors[i-1]
		next := colors[i+1]

		if current == 'A' && prev == 'A' && next == 'A' {
			aliceScore++
		} else if current == 'B' && prev == 'B' && next == 'B' {
			bobScore++
		}
	}
	return aliceScore > bobScore
}
