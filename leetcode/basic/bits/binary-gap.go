package basic

// Consecutive 1s
func BinaryGap(N int) int {
	if N <= 0 || N&(N-1) == 0 {
		return 0
	}
	max := 0
	currCount := -1
	for N > 0 {
		if N&1 > 0 {
			if currCount == -1 || currCount > max {
				max = currCount
			}
			currCount = 0
		} else if currCount != -1 {
			currCount++
		}
		N = N >> 1
	}
	return max
}

// Any 1s
func BinaryGap2(n int) int {
	// TODO

	return 0
}
