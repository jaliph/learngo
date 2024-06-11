package basic

func Rotate(A []int, K int) []int {
	if len(A) == 0 {
		return A
	}

	if K == 0 || K%len(A) == 0 {
		return A
	}

	if K > len(A) {
		K = K % len(A)
	}

	reverse := func(arr *[]int, a, b int) {
		for a < b {
			temp := (*arr)[a]
			(*arr)[a] = (*arr)[b]
			(*arr)[b] = temp
			a++
			b--
		}
	}

	reverse(&A, 0, len(A)-1)
	reverse(&A, 0, K-1)
	reverse(&A, K, len(A)-1)

	return A
}
