package dp

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// state
/**
dp(zero, one, two, three) =
  1 +   // because if the move
	p[zero] * dp(zero, one, two, three) +  if dice has  value where ni dish is left
	p[one] * dp(zero + 1, one - 1, two, three) +
	p[two] * dp(zero, one + 1, two - 1, three) +
	p[three] * dp(zero, one, two + 1, three - 1)


1 - p[zero] * dp(zero, one, two, three) =
	1 +
	p[one] * dp(zero + 1, one - 1, two, three) +
	p[two] * dp(zero, one + 1, two - 1, three) +
	p[three] * dp(zero, one, two + 1, three - 1)

 dp(zero, one, two, three) = (
	1 +
	p[one] * dp(zero + 1, one - 1, two, three) +
	p[two] * dp(zero, one + 1, two - 1, three) +
	p[three] * dp(zero, one, two + 1, three - 1) )   / 1 - p[zero]


	p[zero]  = a / N
	p[one]  = b / N
	p[two]  = c / N
	p[three]  = d / N

dp(zero, one, two, three) = (
	N +
	b * dp(zero + 1, one - 1, two, three) +
	c * dp(zero, one + 1, two - 1, three) +
	d * dp(zero, one, two + 1, three - 1) )   / N - a

**/

func Sushi() {
	Readint := func(in *bufio.Reader) int {
		nStr, _ := in.ReadString('\n')
		nStr = strings.ReplaceAll(nStr, "\r", "")
		nStr = strings.ReplaceAll(nStr, "\n", "")
		n, _ := strconv.Atoi(nStr)
		return n
	}

	ReadLineNumbs := func(in *bufio.Reader) []string {
		line, _ := in.ReadString('\n')
		line = strings.ReplaceAll(line, "\r", "")
		line = strings.ReplaceAll(line, "\n", "")
		numbs := strings.Split(line, " ")
		return numbs
	}

	ReadArrInt := func(in *bufio.Reader) []int {
		numbs := ReadLineNumbs(in)
		arr := make([]int, len(numbs))
		for i, n := range numbs {
			val, _ := strconv.Atoi(n)
			arr[i] = val
		}
		return arr
	}

	in := bufio.NewReader(os.Stdin)

	dishes := Readint(in)
	pieces := ReadArrInt(in)

	cnt := make([]int, 4)
	for _, v := range pieces {
		cnt[v]++
	}

	// fmt.Println(cnt, dishes)

	N := 301
	memoise := make([][][]float64, N)
	for i := 0; i < N; i++ {
		memoise[i] = make([][]float64, N)
		for j := 0; j < N; j++ {
			memoise[i][j] = make([]float64, N)
			for k := 0; k < N; k++ {
				memoise[i][j][k] = float64(-1)
			}
		}
	}

	var solve func(int, int, int, int) float64
	// check the formulae from comments
	solve = func(zero, one, two, three int) float64 {
		if one+two+three == 0 {
			return 0 // no more move left
		}

		if memoise[one][two][three] != -1 {
			return memoise[one][two][three]
		}

		var ans float64 = float64(dishes)
		if one > 0 {
			ans += float64(one) * solve(zero+1, one-1, two, three)
		}
		if two > 0 {
			ans += float64(two) * solve(zero, one+1, two-1, three)
		}
		if three > 0 {
			ans += float64(three) * solve(zero, one, two+1, three-1)
		}
		ans = ans / (float64(dishes - zero))
		memoise[one][two][three] = ans
		return ans
	}

	fmt.Printf("%.10f\n", solve(cnt[0], cnt[1], cnt[2], cnt[3]))
}
