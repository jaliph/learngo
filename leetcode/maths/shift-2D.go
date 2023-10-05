package maths

import "fmt"

func ShiftGrid(grid [][]int, k int) [][]int {
	r, c := len(grid), len(grid[0])

	res := make([][]int, r)
	rows := make([]int, r*c)
	for i := 0; i < r; i++ {
		res[i] = rows[i*c : (i+1)*c]
	}

	for i, _ := range res {
		for j, _ := range res[0] {
			// fmt.Println(_2dTo1d(i, j, c), "->", i, j)
			// fmt.Println(shift1D(_2dTo1d(i, j, c), k, r*c))
			// fmt.Println(_1dtTo2d(shift1D(_2dTo1d(i, j, c), k, r*c), c))
			x, y := _1dtTo2d(shift1D(_2dTo1d(i, j, c), k, r*c), c)
			res[x][y] = grid[i][j]
		}
	}

	return res
}

func Print2D(grid [][]int) {
	for r, _ := range grid {
		for _, colValue := range grid[r] {
			fmt.Print(colValue, " ")
		}
		fmt.Println()
	}
}

func _2dTo1d(i, j, c int) int {
	return i*c + j
}

func shift1D(i, k, c int) int {
	return (i + k) % c
}

func _1dtTo2d(i, c int) (x, y int) {
	x = i / c
	y = i % c
	return
}
