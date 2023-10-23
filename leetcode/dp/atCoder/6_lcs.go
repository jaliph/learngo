package dp

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LCS() {

	in := bufio.NewReader(os.Stdin)

	ReadString := func(in *bufio.Reader) string {
		line, _ := in.ReadString('\n')
		line = strings.ReplaceAll(line, "\r", "")
		line = strings.ReplaceAll(line, "\n", "")
		return line
	}

	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	str1 := ReadString(in)
	str2 := ReadString(in)

	R := len(str1) + 1
	C := len(str2) + 1

	dp := make([][]int, R)
	rows := make([]int, R*C)
	for i := range dp {
		dp[i] = rows[i*C : (i+1)*C]
	}

	for i := 0; i < len(str1); i++ {
		for j := 0; j < len(str2); j++ {
			if str1[i] == str2[j] {
				dp[i+1][j+1] = 1 + dp[i][j]
			} else {
				dp[i+1][j+1] = Max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	i, j := len(str1), len(str2)
	res := ""
	for i > 0 && j > 0 {
		if str1[i-1] == str2[j-1] {
			res = string(str1[i-1]) + res
			i--
			j--
		} else {
			if dp[i-1][j] == dp[i][j] {
				i--
			} else if dp[i][j-1] == dp[i][j] {
				j--
			}
		}
	}

	fmt.Println(res)
}
