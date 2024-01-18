// https://leetcode.com/problems/backspace-string-compare/
package basic

func getNextCharPos(str string, pos int) int {
	bspace := 0
	for pos > -1 {
		if str[pos] == '#' {
			bspace++
		} else if bspace > 0 {
			// delete the corresponding char which is backspace
			bspace--
		} else {
			// now a valid char
			break
		}
		pos--
	}
	return pos
}

func BackspaceCompare(s string, t string) bool {

	sPos := len(s) - 1
	tPos := len(t) - 1

	for sPos > -1 || tPos > -1 {
		sPos = getNextCharPos(s, sPos)
		tPos = getNextCharPos(t, tPos)

		if sPos < 0 && tPos < 0 {
			return true
		}

		if sPos < 0 || tPos < 0 {
			return false
		}

		if s[sPos] != t[tPos] {
			return false
		}
		sPos--
		tPos--
	}

	return true
}

func BackspaceCompare_Stack(s string, t string) bool {
	process := func(str string) string {
		stack := []rune{}
		for _, v := range str {
			if v == '#' && len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else if v != '#' {
				stack = append(stack, v)
			}
		}
		s := ""
		for _, v := range stack {
			s += string(v)
		}
		return s
	}
	return process(s) == process(t)
}
