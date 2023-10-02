package gametheory

import (
	"basic"
	"rwhelper"
)

// for a pile of rocks, you can pick 1 , 2, 3. Find grundy values

func Grundy(n int) int {
	if n == 0 {
		return 0
	}

	set := basic.NewSet()
	for i := n - 1; i >= n-3 && i >= 0; i-- {
		set.Add(Grundy(i))
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

func GrundyRunner() {
	defer rwhelper.CloseWriter()
	n := rwhelper.ReadInt()
	for i := 0; i < n; i++ {
		rwhelper.Write(i, "Grundy is ", Grundy(i))
	}
}
