package graph

import "fmt"

func calculateMinimumHP(dungeon [][]int) int {
	r, c := len(dungeon), len(dungeon[0])
	grid := make([][]int, r)
	rows := make([]int, r*c)

	for i, _ := range grid {
		grid[i] = rows[i*c : (i+1)*c]
	}

	for i := 0; i < c; i++ {
		if i == 0 {
			grid[0][i] = dungeon[0][i]
		} else {
			grid[0][i] = grid[0][i-1] + dungeon[0][i]
		}
	}

	for i := 1; i < r; i++ {
		grid[i][0] = grid[i-1][0] + dungeon[i][0]
	}

	for i := 1; i < r; i++ {
		for j := 1; j < c; j++ {
			grid[i][j] = MaxInt(grid[i-1][j], grid[i][j-1]) + dungeon[i][j]
		}
	}

	Print2D(dungeon)
	Print2D(grid)
	return 0
}

func MaxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Print2D(grid [][]int) {
	for r := range grid {
		for _, colValue := range grid[r] {
			fmt.Print(colValue, " ")
		}
		fmt.Println()
	}
}

func Driver() {
	dungeon := [][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}
	calculateMinimumHP(dungeon)
}
