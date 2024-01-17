package basic

// https://leetcode.com/problems/unique-number-of-occurrences/

func uniqueOccurrences(arr []int) bool {
	freqMap := map[int]int{}
	cMap := map[int]struct{}{}

	for _, v := range arr {
		freqMap[v]++
	}

	for _, v := range freqMap {
		if _, ok := cMap[v]; ok {
			return false
		}
		cMap[v] = struct{}{}
	}

	return true
}
