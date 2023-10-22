// https://atcoder.jp/contests/dp/tasks/dp_e

package dp

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Knapsack_2() {

	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

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

	const P int = 1e5 + 1
	const INF int = 1e12

	dp := make([][]int, Items+1)
	mw := make([][]int, Items+1)

	rows1 := make([]int, (Items+1)*P)
	rows2 := make([]int, (Items+1)*P)

	for i := range rows2 {
		rows2[i] = INF
	}

	for i := range dp {
		dp[i] = rows1[i*P : (i+1)*P]
		mw[i] = rows2[i*P : (i+1)*P]
	}

	/// base
	dp[0][0] = 1
	mw[0][0] = 0

	for val := 0; val < P; val++ {
		for item := 1; item <= Items; item++ {
			// dont take the item
			if dp[item-1][val] == 1 {
				dp[item][val] = 1
				mw[item][val] = Min(mw[item][val], mw[item-1][val])
			}

			// if take the item
			if profits[item-1] <= val && /// if the profit can be taken
				dp[item-1][val-profits[item-1]] == 1 && // if there is 1 from there we will reach here
				mw[item-1][val-profits[item-1]]+weights[item-1] <= Cap { // if taken Cap not violated
				dp[item][val] = 1
				mw[item][val] = Min(mw[item][val], mw[item-1][val-profits[item-1]]+weights[item-1])
			}
		}
	}

	var ans int
	for i := 0; i < P; i++ {
		if dp[Items][i] == 1 {
			ans = i
		}
	}

	fmt.Println(ans)
}
