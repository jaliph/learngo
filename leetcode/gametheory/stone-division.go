// https://www.hackerrank.com/challenges/stone-division/problem?isFullScreen=true

package gametheory

var memoise_SD map[int64]bool

func giveWL(n int64, s []int64) bool {
	if v, ok := memoise_SD[n]; ok {
		return v
	}
	for i := 0; i < len(s); i++ {
		if n%s[i] == 0 && s[i]&1 == 0 {
			memoise_SD[n] = true
			return true
		}
	}
	for i := 0; i < len(s); i++ {
		if n%s[i] == 0 && giveWL(n/s[i], s) == false {
			memoise_SD[n] = true
			return true
		}
	}
	memoise_SD[n] = false
	return false
}

func SstoneDivision(n int64, s []int64) string {
	memoise_SD = make(map[int64]bool)
	if giveWL(n, s) {
		return "First"
	} else {
		return "Second"
	}
}
