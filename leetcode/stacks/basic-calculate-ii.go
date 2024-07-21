package stacks

// https://leetcode.com/problems/basic-calculator-ii/
func Calculate_ii(s string) int {

	var operator byte = '+'
	curr := 0
	prev := 0
	// sign := 1
	// operator := '+'
	res := 0

	operators := map[byte]struct{}{
		'+': {},
		'-': {},
		'/': {},
		'*': {},
	}
	isOperator := func(s byte) bool {
		if _, ok := operators[s]; ok {
			return true
		}
		return false
	}
	isDigit := func(s byte) bool {
		if s >= '0' && s <= '9' {
			return true
		}
		return false
	}

	for i := 0; i < len(s); i++ {
		if isDigit(s[i]) {
			curr = (curr * 10) + int(s[i]-'0')
		}
		if isOperator(s[i]) || (i == len(s)-1) {
			if operator == '+' {
				res += prev
				prev = curr
			} else if operator == '-' {
				res += prev
				prev = -curr
			} else if operator == '*' {
				prev *= curr
			} else if operator == '/' {
				prev /= curr
			}
			operator = s[i]
			curr = 0
		}
	}

	return res + prev
}
