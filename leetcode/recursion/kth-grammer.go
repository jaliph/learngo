package recursion

//https://leetcode.com/problems/k-th-symbol-in-grammar

func KthGrammar(n int, k int) int {
	isSame := true

	N := Pow(2, n)

	for N != 1 {
		N = N / 2

		if k > N {
			k = k - N
			isSame = !isSame
		}
	}

	if isSame {
		return 0
	} else {
		return 1
	}
}

func Pow(a, b int) int {
	res := 1

	for b > 0 {
		if b&1 == 1 {
			res = res * a
		}

		a = a * a
		b = b / 2
	}
	return res
}
