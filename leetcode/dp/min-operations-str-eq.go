// https://leetcode.com/problems/apply-operations-to-make-two-strings-equal/
package dp

func minOperationsDP(s1 string, s2 string, x int) int {
	diff := []int{}
	for i := range s1 {
		if s1[i] != s2[i] {
			diff = append(diff, i)
		}
	}

	if len(diff)&1 == 1 {
		return -1
	}
	if len(diff) == 0 {
		return 0
	}

	dp := make([]float64, len(diff)+1)
	dp[len(diff)] = 0
	dp[len(diff)-1] = float64(x) / float64(2)

	for i := len(diff) - 2; i >= 0; i-- {
		dp[i] = MinF(dp[i+1]+float64(x/2), dp[i+2]+float64(diff[i+1]-diff[i]))
	}

	return int(dp[0])
}

func minOperationsTopDown(s1 string, s2 string, x int) int {
	diff := []int{}
	for i := range s1 {
		if s1[i] != s2[i] {
			diff = append(diff, i)
		}
	}

	if len(diff)&1 == 1 {
		return -1
	}

	l := len(diff)
	memo := make(map[int]float64)
	var recurse func(i int) float64
	recurse = func(i int) float64 {
		if i == l-1 {
			return float64(x) / float64(2)
		} else if i == l {
			return 0
		}

		if v, ok := memo[i]; ok {
			return v
		}

		ans := MinF(
			recurse(i+2)+float64(diff[i+1]-diff[i]),
			recurse(i+1)+(float64(x)/float64(2)))

		memo[i] = ans
		return ans
	}

	return int(recurse(0))
}

func MinF(a float64, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
