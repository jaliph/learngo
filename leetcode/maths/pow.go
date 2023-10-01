package maths

func Pow(x float64, n int) float64 {
	var res float64 = 1
	// power is negative
	if n < 0 {
		n = -n
		x = 1 / x
	}

	for n > 0 {
		if n&1 == 1 {
			res = res * x
		}

		x = x * x
		n /= 2
	}

	return res
}
