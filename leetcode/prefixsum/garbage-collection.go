package prefixsum

// https://leetcode.com/problems/minimum-amount-of-time-to-collect-garbage/

func garbageCollection(garbage []string, travel []int) int {

	prefixSum := make([]int, len(garbage))
	prefixSum[0] = 0
	for i, t := range travel {
		prefixSum[i+1] = prefixSum[i] + t
	}

	last := map[rune]int{
		'P': 0,
		'G': 0,
		'M': 0,
	}
	total := 0
	for T := range last {
		cost := 0
		for i, g := range garbage {
			count := 0
			for _, ch := range g {
				if ch == T {
					count++
				}
			}
			if count > 0 {
				cost += count
				cost += (prefixSum[i] - prefixSum[last[T]])
				last[T] = i
			}
		}
		total += cost
	}

	return total
}

// func Driver() {
// 	garbage := []string{"MMM", "PGM", "GP"}
// 	travel := []int{3, 10}
// 	fmt.Println(garbageCollection(garbage, travel))
// }
