package dp

// https://leetcode.com/problems/wildcard-matching/solutions/3320673/golang-dp-topdown-td-cache-bottomup-2d-array-bu-1d-array/

func IsMatchWP(s string, p string) bool {
	sLen := len(s)
	pLen := len(p)

	dp := map[[2]int]bool{}
	var recur func(int, int) bool
	recur = func(i, j int) bool {
		if i == sLen && j == pLen {
			return true
		}

		if j == pLen {
			return false
		}

		if val, ok := dp[[2]int{i, j}]; ok {
			return val
		}

		if i < sLen && (s[i] == p[j] || p[j] == '?') {
			dp[[2]int{i, j}] = recur(i+1, j+1)
			return dp[[2]int{i, j}]
		}

		if p[j] == '*' {
			res := false
			res = res || recur(i, j+1) // no match
			if i < sLen {
				res = res || recur(i+1, j+1) // match the next
				res = res || recur(i+1, j)   // match the curr
			}
			dp[[2]int{i, j}] = res
			return res
		}
		dp[[2]int{i, j}] = false
		return false
	}

	return recur(0, 0)
}

func IsMatchBU(s string, p string) bool {
	m := len(s)
	n := len(p)

	prev := make([]bool, n+1)
	prev[0] = true

	for j := 0; j < n; j++ {
		if p[j] == '*' {
			prev[j+1] = prev[j]
		} else {
			break
		}
	}

	for i := 0; i < m; i++ {
		curr := make([]bool, n+1)
		for j := 0; j < n; j++ {
			if s[i] == p[j] || p[j] == '?' {
				curr[j+1] = prev[j]
			} else if p[j] == '*' {
				curr[j+1] = prev[j] || prev[j+1] || curr[j]
			}
		}
		prev = curr
	}

	return prev[n]
}
