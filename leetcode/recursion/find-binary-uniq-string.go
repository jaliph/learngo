package recursion

// https://leetcode.com/problems/find-unique-binary-string/

func findDifferentBinaryString(nums []string) string {
	set := map[string]bool{}

	for _, v := range nums {
		set[v] = true
	}

	var solve func(int, string) string

	solve = func(index int, str string) string {
		// base
		if index == len(str) {
			if _, ok := set[str]; ok {
				return ""
			} else {
				return str
			}
		}

		// recur
		str1 := solve(index+1, str+"0")
		if str1 != "" {
			return str1
		}
		str2 := solve(index+1, str+"1")
		if str2 != "" {
			return str2
		}
		return ""
	}

	return solve(0, "")
}
