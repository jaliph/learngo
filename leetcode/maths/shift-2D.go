package maths

import "fmt"

func ShiftGrid(grid [][]int, k int) [][]int {
	r, c := len(grid), len(grid[0])

	res := make([][]int, r)
	rows := make([]int, r*c)
	for i := 0; i < r; i++ {
		res[i] = rows[i*c : (i+1)*c]
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

// func _2dTo1d(i, j int) int {

// }

// func shift1D(i, k int) int {

// }

// func _1dtTo2d(i int) (int, int) {

// }
