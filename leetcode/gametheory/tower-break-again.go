package gametheory

// https://www.hackerrank.com/challenges/tower-breakers-again-1/problem?isFullScreen=true

var div [][]int

const SIZE_DIV = 20 // 1e5

func populateDivisors() {
	div = make([][]int, SIZE_DIV+1)
	for i, _ := range div {
		div[i] = []int{}
	}

	for i := 2; i <= SIZE_DIV; i++ {
		for j := i; j <= SIZE_DIV; j += i {
			div[j] = append(div[j], i)
		}
	}
}

var memoise_grundy_tbr map[int]int

func Grundy_TBR(n int) int {
	if n == 1 {
		return 0
	}
	if v, ok := memoise_grundy_tbr[n]; ok {
		return v
	}
	m := make(map[int]bool)
	for _, v := range div[n] {
		z := n / v
		g := Grundy_TBR(z)
		if v%2 == 0 {
			g = 0
		}
		m[g] = true
	}

	mex := 0
	for i := 0; ; i++ {
		if _, ok := m[i]; !ok {
			mex = i
			break
		}
	}
	memoise_grundy_tbr[n] = mex
	return mex
}

func TowerBreakersAgain(arr []int32) int32 {
	memoise_grundy_tbr = make(map[int]int)
	populateDivisors()
	nim_sum := 0
	for _, v := range arr {
		nim_sum ^= Grundy_TBR(int(v))
	}

	if nim_sum > 0 {
		return 1
	} else {
		return 0
	}
}
