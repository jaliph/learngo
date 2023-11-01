package dp

import "fmt"

// Each vowel 'a' may only be followed by an 'e'.
// Each vowel 'e' may only be followed by an 'a' or an 'i'.
// Each vowel 'i' may not be followed by another 'i'.
// Each vowel 'o' may only be followed by an 'i' or a 'u'.
// Each vowel 'u' may only be followed by an 'a'.

func CountVowelPermutation(n int) int {

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

// Approach 2
const MOD int = 1e9 + 7

func MatrixMultiplication(m1, m2 [][]int) [][]int {
	R1 := len(m1)
	C1 := len(m1[0])
	R2 := len(m2)
	C2 := len(m2[0])

	fmt.Println(m1, m2)
	if C1 != R2 {
		panic("In compatible matrices")
	}

	R := R1
	C := C2
	m3 := make([][]int, R)
	rows := make([]int, R*C)
	for i := range m3 {
		m3[i] = rows[i*C : (i+1)*C]
	}

	for i := 0; i < R1; i++ {
		for j := 0; j < C2; j++ {
			for k := 0; k < R2; k++ {
				m3[i][j] += (m1[i][k] * m2[k][j])
				m3[i][j] %= MOD
			}
		}
	}

	return m3
}

func MatPower(a [][]int, b int) [][]int {
	IdentityMat := [][]int{
		{1, 0, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 0, 0, 0, 1},
	}

	for b > 0 {
		if b&1 == 1 {
			IdentityMat = MatrixMultiplication(IdentityMat, a)
		}
		a = MatrixMultiplication(a, a)
		b = b >> 1
	}

	return IdentityMat
}

func CountVowelPermutation_MatExp(n int) int {
	result := [][]int{{1}, {1}, {1}, {1}, {1}}

	if n >= 2 {
		// Each vowel 'a' may only be followed by an 'e'.
		// Each vowel 'e' may only be followed by an 'a' or an 'i'.
		// Each vowel 'i' may not be followed by another 'i'.
		// Each vowel 'o' may only be followed by an 'i' or a 'u'.
		// Each vowel 'u' may only be followed by an 'a'.
		T := [][]int{
			{0, 1, 0, 0, 0},
			{1, 0, 1, 0, 0},
			{1, 1, 0, 1, 1},
			{0, 0, 1, 0, 1},
			{1, 0, 0, 0, 0},
		}
		result = MatrixMultiplication(MatPower(T, n-1), result)
	}
	fmt.Println(result)

	sum := 0
	for i := range result {
		sum += result[i][0]
		sum += MOD
	}
	return sum
}

// func Driver() {
// 	fmt.Println(CountVowelPermutation_MatExp(5))
// }
