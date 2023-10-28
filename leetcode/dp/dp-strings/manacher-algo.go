package dp

import (
	"fmt"
)

type Manacher struct {
	p         []int  // manacher array
	s         string // original string
	t         string // processed string with # in between..
	tlen      int
	lpsStart  int
	lpsLength int
}

func NewManacher(str string) *Manacher {
	s := Manacher{}
	s.s = str

	t := "@"
	for _, v := range str {
		t += "#" + string(v)
	}
	t += "#$"

	s.t = t
	s.tlen = len(s.t)
	s.p = make([]int, s.tlen)
	return &s
}

func (m *Manacher) RunManacher() {
	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	c, r := 0, 0
	for i := 1; i < m.tlen-1; i++ {
		if i < r {
			m.p[i] = Min(r-i, m.p[2*c-i])
		}

		// expand the palindromes at this place
		for m.t[i-m.p[i]-1] == m.t[i+m.p[i]+1] {
			m.p[i]++
		}

		if i+m.p[i] > r {
			r = i + m.p[i]
			c = i
		}

		if m.p[i] > m.lpsLength {
			m.lpsStart = (i - m.p[i] - 1) / 2
			m.lpsLength = m.p[i]
		}
	}
}

func Driver() {
	s := "abbd"

	m := NewManacher(s)
	m.RunManacher()
	fmt.Println(m)

	l, r := m.lpsStart, m.lpsStart+m.lpsLength-1
	fmt.Println(s[l : r+1])
}

// 0, 0, 1, 0, 1, 2, 1, 0, 1, 0, 1, 2, 1, 0, 7, 0, 1, 2, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0
// 0, 1, 2, 1, 2, 3, 2, 1, 2, 1, 2, 3, 2, 1, 8, 1, 2, 3, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 0

/**
type Manacher struct {
	p         []int  // manacher array
	s         string // original string
	t         string // processed string with # in between..
	tlen      int
	lpsCenter int
	lpsLength int
}

func NewManacher(str string) *Manacher {
	s := Manacher{}
	s.s = str

	t := ""
	for _, v := range str {
		t += "#" + string(v)
	}
	t += "#"

	s.t = t
	s.tlen = len(s.t)
	s.p = make([]int, s.tlen)
	for i := range s.p {
		s.p[i] = 1
	}

	return &s
}

func (m *Manacher) RunManacher() {
	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	// p[i] = defines how many length of palindrome we have at index i.
	// Actual length of palindrome at p[i] is p[i] - 1, as We initially assigned 1 value to every p[i], So '#' characters also got length 1, which is why we have to subtract 1.
	l, r := 1, 1
	for i := 1; i < m.tlen; i++ {
		// update p[i]
		if l+r-i >= 0 {
			m.p[i] = Max(0, Min(r-i, m.p[l+r-i]))
		}

		// expand the palindromes at this place
		for i-m.p[i] >= 0 && i+m.p[i] < m.tlen && m.t[i-m.p[i]] == m.t[i+m.p[i]] {
			m.p[i]++
		}

		// update l and r
		if i+m.p[i] > r {
			l = i - m.p[i]
			r = i + m.p[i]

			if m.p[i] > m.lpsLength {
				m.lpsLength = m.p[i] - 1
				m.lpsCenter = i
			}
		}
	}

	// fmt.Println(m.p)
}

func (m *Manacher) GetLongest(center int, isOdd bool) int {
	var pos int
	// if it is even, count the hash in
	if isOdd {
		pos = 2*center + 1
	} else {
		pos = 2*center + 1 + 1
	}

	return m.p[pos] - 1
}

func (m *Manacher) IsPalindrome(l, r int) bool {
	// (r - l + 1) is the size of the current window, if current window size comes under its center's biggest palindrome, then it means it's also palindrome, as part of palindrome is also palindrome.

	// r%2 == l%2 tells if r and l forms length of odd length
	if r-l+1 <= m.GetLongest((r+l)/2, r%2 == l%2) {
		return true
	} else {
		return false
	}
}


func Driver() {
	s := "abbd"

	m := NewManacher(s)
	m.RunManacher()
	fmt.Println(strings.ReplaceAll(m.t[m.lpsCenter-m.lpsLength:m.lpsCenter+m.lpsLength], "#", ""))
	// fmt.Println(ans)
}

**/
