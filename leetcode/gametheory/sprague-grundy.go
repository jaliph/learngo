package gametheory

import (
	"basic"
	"fmt"
)

// for a pile of rocks, you can pick 1, 2, 3. Find grundy values

var memoise map[int]int

func Grundy(n int) int {
	if n == 0 {
		return 0
	}
	if v, ok := memoise[n]; ok {
		return v
	}

	set := basic.NewSet()
	for i := n - 1; i >= n-3 && i >= 0; i-- {
		set.Add(Grundy(i))
	}

	fmt.Println(n, " has Grundy's", set)
	mex := 0
	for {
		if set.Has(mex) {
			mex++
		} else {
			memoise[n] = mex
			return mex
		}
	}
}

func GrundyRunner() {
	memoise = map[int]int{}
	for i := 0; i < 9; i++ {
		fmt.Println(i, "Grundy is ", Grundy(i))
	}
}

// 0 1 1 1 0 1 1 1 0

// grundy(x) = Mex(grundy(y1), grundy(y2), grundy(y2))
