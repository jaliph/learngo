package prefixsum

import "fmt"

// https://leetcode.com/problems/number-of-submatrices-that-sum-to-target/

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	R := len(matrix)
	C := len(matrix[0])

	for i := 0; i < R; i++ {
		for j := 1; j < C; j++ {
			matrix[i][j] = matrix[i][j-1] + matrix[i][j]
		}
	}

	count := 0

	for i := 0; i < C; i++ {
		for j := i; j < C; j++ {
			mp := map[int]int{}
			mp[0] = 1
			sum := 0
			for k := 0; k < R; k++ {
				sum += matrix[k][j]
				if i > 0 {
					sum -= matrix[k][i-1]
				}
				if v, ok := mp[sum-target]; ok {
					count += v
				}
				mp[sum]++
			}

		}
	}
	return count
}

func Driver() {
	matrix := [][]int{
		{0, 1, 0},
		{1, 1, 1},
		{0, 1, 0},
	}
	target := 0

	fmt.Println(numSubmatrixSumTarget(matrix, target))
}
