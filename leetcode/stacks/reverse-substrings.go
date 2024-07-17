package stacks

func ReverseParentheses(s string) string {
	stack := ""
	for _, val := range s {
		if val == ')' {
			temp := ""
			i := len(stack) - 1
			for stack[i] != '(' {
				temp += string(stack[i])
				i--
			}
			stack = stack[:i] + temp
		} else {
			stack += string(val)
		}
	}
	return stack

	// stack := []string{""}

	// reverse := func(s string) string {
	// 	res := ""
	// 	for _, v := range s {
	// 		res = string(v) + res
	// 	}
	// 	fmt.Println("rev", s, res)
	// 	return res
	// }

	// for _, v := range s {
	// 	fmt.Println(stack, v)
	// 	if v == '(' {
	// 		stack = append(stack, "")
	// 	} else if v != ')' {
	// 		stack[len(stack)-1] += string(v)
	// 	} else {
	// 		top := stack[len(stack)-1]
	// 		stack = stack[:len(stack)-1]
	// 		stack[len(stack)-1] += reverse(top)
	// 	}
	// }
	// return stack[0]
}
