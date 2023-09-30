package basic

// https://leetcode.com/problems/valid-anagram/
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := make(map[rune]int)

	for _, v := range s {
		sMap[v]++
	}

	tMap := make(map[rune]int)

	for _, v := range t {
		tMap[v]++
	}

	if len(sMap) != len(tMap) {
		return false
	}

	for k, v1 := range sMap {
		if v2, ok := tMap[k]; ok && (v1 == v2) {
			continue
		} else {
			return false
		}
	}

	return true
}

func IsAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := make(map[rune]int)

	for _, v := range s {
		sMap[v]++
	}

	for _, c := range t {
		if count, exists := sMap[c]; !exists || count == 0 {
			return false
		}
		sMap[c]--
	}

	return true
}
