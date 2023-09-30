package basic

// https://leetcode.com/problems/isomorphic-strings/
func IsIsomorphic(s string, t string) bool {
	sMap := [256]byte{}
	tMap := [256]byte{}

	for idx, v := range s {
		if sMap[v] != tMap[t[idx]] {
			return false
		}
		sMap[v] = byte(idx + 10)
		tMap[t[idx]] = byte(idx + 10)
	}

	return true
}

func IsIsomorphicBrute(s string, t string) bool {
	sMap := make(map[rune]int)
	tMap := make(map[rune]int)

	t_rune := []rune(t)
	for idx, v := range s {
		if sMap[v] != tMap[t_rune[idx]] {
			return false
		}
		sMap[v] = idx + 10
		tMap[t_rune[idx]] = idx + 10
	}

	return true
}
