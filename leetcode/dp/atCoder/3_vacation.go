// https://atcoder.jp/contests/dp/tasks/dp_c
package dp

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Vacations() {

	in := bufio.NewReader(os.Stdin)

	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

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

	days := Readint(in)
	var prev []int
	for d := 0; d < days; d++ {
		if d == 0 {
			prev = ReadArrInt(in)
		} else {
			activities := ReadArrInt(in)
			activities[0] += Max(prev[1], prev[2])
			activities[1] += Max(prev[0], prev[2])
			activities[2] += Max(prev[0], prev[1])

			prev = activities
		}
		// fmt.Println(prev)
	}
	fmt.Println(Max(Max(prev[0], prev[1]), prev[2]))
}
