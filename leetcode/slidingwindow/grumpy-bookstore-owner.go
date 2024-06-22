package slidingwindow

// https://leetcode.com/problems/grumpy-bookstore-owner/

func MaxSatisfied(customers []int, grumpy []int, minutes int) int {

	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// total how much i can be satisfied
	total := 0
	for i := range grumpy {
		if grumpy[i] == 0 {
			total += customers[i]
		}
	}
	// no matter how much minutes i have this is atleast satisfied

	// lets search the window of minutes
	// and find out max of grumpy entries

	wStart := 0
	count := 0
	max := 0
	for wEnd := 0; wEnd < len(grumpy); wEnd++ {
		if grumpy[wEnd] == 1 {
			count += customers[wEnd]
		}

		if wEnd > minutes-1 {
			if grumpy[wStart] == 1 {
				count -= customers[wStart]
			}
			wStart++
		}
		max = Max(max, count)
	}

	return total + max
}
