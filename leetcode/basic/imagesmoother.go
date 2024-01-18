// https://leetcode.com/problems/image-smoother

package basic

func imageSmoother(img [][]int) [][]int {
	R := len(img)
	C := len(img[0])

	ans := make([][]int, R)
	rows := make([]int, R*C)
	for r := range ans {
		ans[r] = rows[r*C : (r+1)*C]
	}

	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}
	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {

			count := 0
			sum := 0

			for ii := Max(0, i-1); ii < Min(R, i+2); ii++ {
				for jj := Max(0, j-1); jj < Min(C, j+2); jj++ {

					count++
					sum += img[ii][jj]
				}
			}

			ans[i][j] = sum / count
		}
	}

	return ans
}
