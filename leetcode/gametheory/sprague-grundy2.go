package gametheory

import (
	"basic"
	"rwhelper"
)

// for a pile of rocks, you can pick any n which is divisible by n including 1, & n Find grundy values

func Grundy2(n int) int {
	if n == 0 {
		return 0
	}

	set := basic.NewSet()
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			set.Add(Grundy2(n - i))
		}
	}

	mex := 0
	for {
		if set.Has(mex) {
			mex++
		} else {
			return mex
		}
	}
}

func Grundy2Runner() {
	defer rwhelper.CloseWriter()
	n := rwhelper.ReadInt()
	for i := 0; i < n; i++ {
		rwhelper.Write(i, "Grundy is ", Grundy2(i))
	}
}
