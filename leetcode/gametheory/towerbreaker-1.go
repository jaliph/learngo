// https://www.hackerrank.com/challenges/tower-breakers-1

package gametheory

func TowerBreakers(n int32, m int32) int32 {
	if m == 1 || n%2 == 0 {
		return 2
	} else {
		return 1
	}
}

// nim sum
// if m == 1 , first player cannot move. P Position
// if n & 1 == 0, P position. what ever 1st player does, second player can balance out
// if n is odd, N position, first player wins
