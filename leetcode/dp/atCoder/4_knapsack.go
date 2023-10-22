// https://atcoder.jp/contests/dp/tasks/dp_d
package dp

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Knapsack_1() {

	// Input Helpers
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
	initial := ReadArrInt(in)

	Cap := initial[1]
	Items := initial[0]
	weights := []int{}
	profits := []int{}
	for i := 0; i < Items; i++ {
		items := ReadArrInt(in)
		weights = append(weights, items[0])
		profits = append(profits, items[1])
	}

	dp := make([]int, Cap+1)
	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// max(solve(i + 1, cap), profit[i] + solve(i + 1, cap - weight[i]))
	for w := 0; w < Items; w++ {
		for c := Cap; c >= 1; c-- {
			if weights[w] <= c {
				dp[c] = Max(dp[c], profits[w]+dp[c-weights[w]])
			}
		}
	}

	fmt.Println(dp[Cap])
}
