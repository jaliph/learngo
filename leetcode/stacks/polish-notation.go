package stacks

import "strconv"

func evalRPN(tokens []string) int {
	stack := []int{}
	operators := map[string]struct{}{
		"+": {},
		"-": {},
		"*": {},
		"/": {},
	}

	var result int
	for _, v := range tokens {
		if _, ok := operators[v]; !ok {
			i, _ := strconv.Atoi(v)
			stack = append(stack, i)
		} else {
			switch v {
			case "+":
				result = stack[len(stack)-2] + stack[len(stack)-1]
			case "-":
				result = stack[len(stack)-2] - stack[len(stack)-1]
			case "*":
				result = stack[len(stack)-2] * stack[len(stack)-1]
			case "/":
				result = stack[len(stack)-2] / stack[len(stack)-1]
			}
			stack = stack[:len(stack)-2]
			stack = append(stack, result)
		}
	}
	return stack[0]
}
