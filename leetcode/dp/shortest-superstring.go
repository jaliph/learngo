package dp

type shortestStringState struct {
	str string
	m   int
}

// DP
func shortestSuperstring_DP(words []string) string {
	mask := 1<<len(words) - 1

	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	overLapFinder := func(str1, str2 string) string {
		str := str1 + str2
		var i int
		for i = 0; i < len(str1); i++ {
			if str1[i:] == str2[0:Min(len(str1)-i, len(str2))] {
				str = str1 + str2[len(str1)-i:]
				break
			}
		}
		return str
	}

	dp := map[shortestStringState]string{}
	var solve func(string, int) string
	solve = func(word string, mask int) string {
		// base case

		if mask == 0 {
			return word
		}

		if v, ok := dp[shortestStringState{word, mask}]; ok {
			return v
		}

		var str string
		for i, w := range words {
			var superString string
			if ((mask >> i) & 1) > 0 {
				superString = solve(w, mask^(1<<i))

				overLappedString := overLapFinder(word, superString)
				if len(str) == 0 || len(overLappedString) < len(str) {
					str = overLappedString
				}
			}
		}

		dp[shortestStringState{word, mask}] = str
		return str
	}

	return solve("", mask)
}
