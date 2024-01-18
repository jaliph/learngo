package basic

import "fmt"

type Random struct {
	a   int
	b   int
	x0  int // seed
	mod int
}

func GetRandomGen(a, b, x0, mod int) *Random {
	return &Random{
		a, b, x0, mod,
	}
}

func (r *Random) RandInt() int {
	r.x0 = ((r.a * r.x0) + r.b) % r.mod
	return r.x0
}

func (r *Random) RandBinarySequence(length int) string {
	s := ""
	for i := 0; i < length; i++ {
		r.x0 = ((r.a * r.x0) + r.b) % r.mod
		s += fmt.Sprintf("%d", r.x0%2)
	}
	return s
}
