// https://atcoder.jp/contests/dp/tasks/dp_k

package dp

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func StoneGame() {
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

	line1 := ReadArrInt(in)
	stones := ReadArrInt(in)
	pile := line1[1]

	dp := make([]int, pile+1)

	dp[0] = 0 // obvio in golang
	for i := 1; i <= pile; i++ {
		for j := 0; j < len(stones); j++ {
			if i-stones[j] >= 0 && dp[i-stones[j]] == 0 {
				dp[i] = 1
				break
			}
		}
	}

	if dp[pile] == 1 {
		fmt.Println("First")
	} else {
		fmt.Println("Second")
	}
}
