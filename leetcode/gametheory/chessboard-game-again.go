package gametheory

import "fmt"

var dx = []int{-2, -2, 1, -1}
var dy = []int{1, -1, -2, -2}

var memoise_cba map[string]int

func Key(x, y int) string {
	return fmt.Sprintf("%d#%d", x, y)
}

func Grundy_CBA(x, y int) int {
	if v, ok := memoise_cba[Key(x, y)]; ok {
		return v
	}
	m := make(map[int]bool)
	for i := 0; i < 4; i++ {
		nx := x + dx[i]
		ny := y + dy[i]
		if nx >= 1 && ny >= 1 && nx <= 15 && ny <= 15 {
			m[Grundy_CBA(nx, ny)] = true
		}
	}

	mex := 0
	for i := 0; ; i++ {
		if _, ok := m[i]; !ok {
			mex = i
			break
		}
	}
	memoise_cba[Key(x, y)] = mex
	return mex
}

// https://www.hackerrank.com/challenges/chessboard-game-again-1
func ChessboardGameAgain(coins [][]int32) string {
	memoise_cba = make(map[string]int)
	var nim_sum int = 0
	for _, v := range coins {
		nim_sum ^= Grundy_CBA(int(v[0]), int(v[1]))
	}

	if nim_sum > 0 {
		return "First"
	} else {
		return "Second"
	}
}

/**
first positive integer
Grundy(x) = Mex(Grundy(y1), Grundy(y2), Grundy(y3))

                  0             3         4

									1


0 1 3 5

2

**/
