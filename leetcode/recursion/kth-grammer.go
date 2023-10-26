package recursion

//https://leetcode.com/problems/k-th-symbol-in-grammar

func KthGrammar(n int, k int) int {
	// result is 0 in the base case
	res := 0

	// 1st position will always be zero
	if k == 0 {
		return 0
	}

	l := 0
	r := Pow(2, n-1)

	var mid int
	for i := 1; i < n; i++ {
		mid = l + (r-l)/2
		// at every level, check which segment k belongs to
		if k <= mid {
			r = mid
		} else {
			l = mid + 1

			// k belongs to the second half, reverse it
			if res == 1 {
				res = 0
			} else {
				res = 1
			}
		}
	}
	return res
}

func KthGrammar_1(n int, k int) int {
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
