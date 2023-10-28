package dp

const MOD int = 1e9 + 7

// Each vowel 'a' may only be followed by an 'e'.
// Each vowel 'e' may only be followed by an 'a' or an 'i'.
// Each vowel 'i' may not be followed by another 'i'.
// Each vowel 'o' may only be followed by an 'i' or a 'u'.
// Each vowel 'u' may only be followed by an 'a'.

func countVowelPermutation(n int) int {

	a, e, i, o, u := 0, 1, 2, 3, 4
	prev := make([]int, 5)
	for i := range prev {
		prev[i] = 1
	}

	for j := 2; j <= n; j++ {
		dp := make([]int, 5)
		dp[a] = AddVModulo(dp[a], prev[e])
		dp[e] = AddVModulo(dp[e], prev[a], prev[i])
		dp[i] = AddVModulo(dp[i], prev[a], prev[e], prev[o], prev[u])
		dp[o] = AddVModulo(dp[o], prev[i], prev[u])
		dp[u] = AddVModulo(dp[u], prev[a])
		prev = dp
	}

	return AddVModulo(prev...)
}

func AddVModulo(args ...int) int {
	sum := 0
	for _, v := range args {
		sum += v
		sum %= MOD
	}
	return sum
}
