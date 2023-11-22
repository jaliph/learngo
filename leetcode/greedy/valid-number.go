package greedy

import (
	"strings"
)

// https://leetcode.com/problems/valid-number

func isNumber(s string) bool {
	s = strings.Trim(s, " ")

	pointSeen := false
	numberSeen := false
	eSeen := false
	numberSeenAfterE := true
	for i, ch := range s {
		if ch <= '9' && ch >= '0' {
			numberSeen = true
			numberSeenAfterE = true
		} else if ch == '+' || ch == '-' {
			if !(i == 0 || (i > 0 && (s[i-1] == 'e' || s[i-1] == 'E'))) {
				return false
			}
		} else if ch == '.' {
			if eSeen || pointSeen {
				return false
			}
			pointSeen = true
		} else if ch == 'E' || ch == 'e' {
			if eSeen || !numberSeen {
				return false
			}
			eSeen = true
			numberSeenAfterE = false
		} else {
			return false
		}
	}

	return numberSeen && numberSeenAfterE
}

func Driver() {
	isNumber("89977")
}
