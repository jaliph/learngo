package basic

// https://leetcode.com/problems/find-the-winner-of-the-circular-game
func FindTheWinner(n int, k int) int {
	circle := make([]int, n)
	for i := range circle {
		circle[i] = i + 1
	}

	pos := 0
	for len(circle) != 1 {
		pos = (pos + (k - 1)) % len(circle)
		circle = append(circle[:pos], circle[pos+1:]...)
	}
	return circle[0]
}
