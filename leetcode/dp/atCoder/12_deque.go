// https: //atcoder.jp/contests/dp/tasks/dp_l

package dp

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// TLE
func Deque() {
	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
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

	ReadInt := func(in *bufio.Reader) int {
		line, _ := in.ReadString('\n')
		line = strings.ReplaceAll(line, "\r", "")
		line = strings.ReplaceAll(line, "\n", "")
		val, _ := strconv.Atoi(line)
		return val
	}

	in := bufio.NewReader(os.Stdin)

	cnt := ReadInt(in)
	vals := ReadArrInt(in)

	dp := make([][]int, cnt)
	rows := make([]int, cnt*cnt)
	for i := range rows {
		rows[i] = -1
	}

	for i := range dp {
		dp[i] = rows[i*cnt : (i+1)*cnt]
	}

	var solve func(int, int) int
	solve = func(l, r int) int {
		// base
		if l > r {
			return 0
		}

		if l == r {
			return vals[l]
		}

		if dp[l][r] != -1 {
			return dp[l][r]
		}

		res := Max(vals[l]-solve(l+1, r), vals[r]-solve(l, r-1))

		dp[l][r] = res
		return res
	}

	sum := 0
	for _, v := range vals {
		sum += v
	}

	fmt.Println(solve(0, cnt-1))
}

// TLE
func Deque_1() {
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

	ReadInt := func(in *bufio.Reader) int {
		line, _ := in.ReadString('\n')
		line = strings.ReplaceAll(line, "\r", "")
		line = strings.ReplaceAll(line, "\n", "")
		val, _ := strconv.Atoi(line)
		return val
	}

	in := bufio.NewReader(os.Stdin)

	cnt := ReadInt(in)
	vals := ReadArrInt(in)

	dp := map[[3]int]int{}
	var solve func(int, int, int) int
	solve = func(l, r int, isTaro int) int {
		// base
		if l > r {
			return 0
		}

		if v, ok := dp[[3]int{l, r, isTaro}]; ok {
			return v
		}

		var res int
		if isTaro == 1 {
			res = Max(vals[l]+solve(l+1, r, isTaro*-1), vals[r]+solve(l, r-1, isTaro*-1))
		} else {
			res = Min(solve(l+1, r, isTaro*-1), solve(l, r-1, isTaro*-1))
		}

		dp[[3]int{l, r, isTaro}] = res
		return res
	}

	sum := 0
	for _, v := range vals {
		sum += v
	}

	X := solve(0, cnt-1, 1)
	Y := sum - X

	fmt.Println(X - Y)
}
