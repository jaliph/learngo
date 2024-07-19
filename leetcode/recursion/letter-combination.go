package recursion

// https://leetcode.com/problems/letter-combinations-of-a-phone-number/
func LetterCombinations(digits string) []string {
	t9Map := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}

	res := []string{""}

	for _, d := range digits {
		currLength := len(res)
		modArr := []string{}
		for i := 0; i < currLength; i++ {
			currElem := res[i]
			for _, ch := range t9Map[string(d)] {
				modArr = append(modArr, currElem+ch)
			}
		}
		res = modArr
	}

	return res
}
