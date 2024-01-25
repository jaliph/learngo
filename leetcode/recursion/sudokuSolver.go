// https://leetcode.com/problems/sudoku-solver/

package recursion

func solveSudoku(board [][]byte) {

	canFillValue := func(i, j int, ans byte) bool {
		for x := 0; x < 9; x++ {
			if board[x][j] == ans {
				return false
			}
		}

		for y := 0; y < 9; y++ {
			if board[i][y] == ans {
				return false
			}
		}

		// check the square
		sx := (i / 3) * 3
		sy := (j / 3) * 3

		for x := sx; x < sx+3; x++ {
			for y := sy; y < sy+3; y++ {
				if board[x][y] == ans {
					return false
				}
			}
		}

		return true
	}

	var solve func(i, j int) bool
	solve = func(i, j int) bool {
		// base
		if i == 9 {
			return true
		}

		if j == 9 {
			return solve(i+1, 0)
		}

		if board[i][j] != '.' {
			return solve(i, j+1)
		}

		// recur
		for ans := '1'; ans <= '9'; ans++ {
			if canFillValue(i, j, byte(ans)) {
				board[i][j] = byte(ans)
				if solve(i, j+1) {
					return true
				}
				board[i][j] = '.'
			}
		}
		return false
	}

	solve(0, 0)
}

// func Driver() {
// 	board := [][]byte{
// 		{'.', '.', '.', '1', '5', '.', '.', '.', '6'},
// 		{'.', '.', '.', '.', '.', '8', '.', '3', '.'},
// 		{'5', '.', '.', '.', '.', '.', '4', '.', '.'},
// 		{'.', '1', '.', '.', '.', '.', '.', '.', '4'},
// 		{'.', '.', '.', '.', '.', '7', '.', '.', '.'},
// 		{'.', '.', '6', '4', '.', '.', '.', '7', '.'},
// 		{'.', '.', '.', '.', '8', '.', '.', '2', '.'},
// 		{'4', '9', '.', '.', '1', '.', '.', '.', '3'},
// 		{'.', '.', '7', '2', '6', '.', '.', '.', '1'}}
// 	Print2D := func(grid [][]byte) {
// 		for r := range grid {
// 			for _, colValue := range grid[r] {
// 				val := string(colValue)
// 				fmt.Print(val, " ")
// 			}
// 			fmt.Println()
// 		}
// 		fmt.Println("----")
// 	}

// 	Print2D(board)

// 	solveSudoku(board)
// 	Print2D(board)
// }
