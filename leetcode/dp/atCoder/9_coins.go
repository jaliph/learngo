// https://atcoder.jp/contests/dp/tasks/dp_i

package dp

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Coins() {

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

	ReadArrFloat64 := func(in *bufio.Reader) []float64 {
		numbs := ReadLineNumbs(in)
		arr := make([]float64, len(numbs))
		for i, n := range numbs {
			val, _ := strconv.ParseFloat(n, 64)
			arr[i] = val
		}
		return arr
	}

	in := bufio.NewReader(os.Stdin)

	coins := Readint(in)
	probabs := ReadArrFloat64(in)

	memoise := map[[2]int]float64{}

	var calculateProbability func(int, int) float64
	calculateProbability = func(i, heads int) float64 {
		tails := coins - i - heads
		// base
		if i == coins {
			return 1
		}

		if v, ok := memoise[[2]int{i, heads}]; ok {
			return v
		}

		var ans float64 = 0
		if heads > 0 {
			ans += probabs[i] * calculateProbability(i+1, heads-1)
		}
		if tails > 0 {
			ans += (1 - probabs[i]) * calculateProbability(i+1, heads)
		}

		memoise[[2]int{i, heads}] = ans
		return ans
	}

	var ans float64 = 0
	for i := (coins / 2) + 1; i <= coins; i++ {
		ans += calculateProbability(0, i)
	}
	fmt.Printf("%.10f\n", ans)
}
