package basic

func RotateMatrix(matrix [][]int) {

	l := 0
	r := len(matrix)

	for l < r {
		for i := 0; i < r-l+1; i++ {
			top := l
			bottom := r

			temp := matrix[top][l+i]

			matrix[top][l+i] = matrix[bottom-i][l]

			matrix[bottom-1][l] = matrix[bottom][r-i]

			matrix[bottom][r-i] = matrix[top+i][r]

			matrix[top+i][r] = temp

		}
		l++
		r--
	}
}
