package dp

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://atcoder.jp/contests/dp/tasks/dp_h

func GridGame_1() {
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

	ReadGrid := func(in *bufio.Reader) []string {
		line, _ := in.ReadString('\n')
		line = strings.ReplaceAll(line, "\r", "")
		line = strings.ReplaceAll(line, "\n", "")
		numbs := strings.Split(line, "")
		return numbs
	}

	in := bufio.NewReader(os.Stdin)

	line1 := ReadArrInt(in)
	R := line1[0]
	C := line1[1]
	grid := make([][]string, R)

	for i := 0; i < R; i++ {
		grid[i] = ReadGrid(in)
	}

	const MOD int = 1e9 + 7
	prev := make([]int, C)

	prev[0] = 1

	for j := 1; j < C; j++ {
		if grid[0][j] == "." {
			prev[j] = prev[j-1]
		}
	}

	fmt.Println(prev)

	for i := 1; i < R; i++ {
		dp := make([]int, C)
		if prev[0] == 1 && grid[i][0] == "." {
			dp[0] = 1
		}
		for j := 1; j < C; j++ {
			if grid[i][j] == "." {
				dp[j] = prev[j] + dp[j-1]
				dp[j] %= MOD
			}
		}
		fmt.Println(dp)
		prev = dp
	}

	fmt.Println(prev[C-1])
}
