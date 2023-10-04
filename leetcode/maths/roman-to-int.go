package maths

//https://leetcode.com/problems/roman-to-integer/

func RomanToInt(s string) int {
	rMap := make(map[byte]int)

	rMap['I'] = 1
	rMap['V'] = 5
	rMap['X'] = 10
	rMap['L'] = 50
	rMap['C'] = 100
	rMap['D'] = 500
	rMap['M'] = 1000

	val := 0
	for idx, ch := range s {
		if idx+1 < len(s) && rMap[byte(ch)] < rMap[s[idx+1]] {
			val -= rMap[byte(ch)]
		} else {
			val += rMap[byte(ch)]
		}
	}
	return val
}
