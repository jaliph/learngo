package dp

var dp [][]int
var Max func(int, int) int

func LongestPalindrome(s string) string {

	str := ""
	max := 0

	Max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	N := len(s)
	dp = make([][]int, N)
	for i := range dp {
		dp[i] = make([]int, N)
	}

	for i := range dp {
		dp[i][i] = 1
	}

	findPalindrome_BU(s)

	for i := N - 1; i >= 0; i-- {
		for j := i; j < N; j++ {
			if dp[i][j] > max {
				max = dp[i][j]
				str = s[i : j+1]
			}
		}
	}

	return str
}

func findPalindrome_BU(s string) {
	N := len(s)
	for i := N - 1; i >= 0; i-- {
		for j := i + 1; j < N; j++ {
			if s[i] == s[j] {
				if j-i+1 == 2+dp[i+1][j-1] {
					dp[i][j] = 2 + dp[i+1][j-1]
					continue
				}
			}
			dp[i][j] = Max(dp[i+1][j], dp[i][j-1])
		}
	}
}

func findPalindrome_TD(l, r int, s string) int {
	if l > r {
		return 0
	}

	if dp[l][r] != 0 {
		return dp[l][r]
	}

	if s[l] == s[r] {
		sublength := 2 + findPalindrome_TD(l+1, r-1, s)
		if sublength == r-l+1 {
			dp[l][r] = sublength
			return sublength
		}
	}

	res := Max(findPalindrome_TD(l+1, r, s), findPalindrome_TD(l, r-1, s))
	dp[l][r] = res
	return res
}

// func Driver() {
// 	fmt.Println(LongestPalindrome("cbbd"))
// }
