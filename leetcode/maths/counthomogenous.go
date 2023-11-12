package maths

func countHomogenous(s string) int {
	left := 0
	res := 0

	const MOD int = 1e9 + 7
	for right := range s {
		if s[left] == s[right] {
			res += right - left + 1
		} else {
			left = right
			res += 1
		}
		res = res % MOD
	}
	return res
}

// func Driver() {
// 	fmt.Println(countHomogenous("sss"))
// }
