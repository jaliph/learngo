package stacks

func Calculate(s string) int {

	var eval func(idx int, s string) (int, int)
	eval = func(idx int, s string) (int, int) {
		res := 0
		sign := 1
		curr := 0
		var j int
		for j = idx; j < len(s); j++ {
			if s[j] >= '0' && s[j] <= '9' {
				curr = (curr * 10) + int(s[j]-'0')
			} else if s[j] == '(' {
				curr, j = eval(j+1, s)
			} else if s[j] == '+' || s[j] == '-' {
				res += (curr * sign)

				curr = 0

				if s[j] == '-' {
					sign = -1
				} else {
					sign = 1
				}
			} else if s[j] == ')' {
				break
			}
		}
		return res + (curr * sign), j
	}
	res, _ := eval(0, s)
	return res
}
