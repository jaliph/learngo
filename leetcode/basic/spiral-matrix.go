package basic

// https://leetcode.com/problems/spiral-matrix/

func spiralOrder(matrix [][]int) []int {

	res := []int{}
	l := 0
	r := len(matrix[0])
	top := 0
	bottom := len(matrix)

	for l < r && top < bottom {

		// top
		for i := l; i < r; i++ {
			res = append(res, matrix[top][i])
		}
		top++

		// right
		for i := top; i < bottom; i++ {
			res = append(res, matrix[i][r-1])
		}
		r--

		if !(l < r && top < bottom) {
			break
		}

		// bottom
		for i := r - 1; i >= l; i-- {
			res = append(res, matrix[bottom-1][i])
		}
		bottom--

		// left
		for i := bottom - 1; i >= top; i-- {
			res = append(res, matrix[i][l])
		}
		l++
	}
	return res
}
