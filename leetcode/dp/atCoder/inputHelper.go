package dp

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Input_Optimised() {

	// frog_2.go
	in := bufio.NewReader(os.Stdin)
	testDetails := ReadArrInt(in)
	stones := ReadArrInt(in)
	res := Frog_2(testDetails[0], testDetails[1], stones)
	fmt.Println(res)
}

func Readint(in *bufio.Reader) int {
	nStr, _ := in.ReadString('\n')
	nStr = strings.ReplaceAll(nStr, "\r", "")
	nStr = strings.ReplaceAll(nStr, "\n", "")
	n, _ := strconv.Atoi(nStr)
	return n
}

func ReadLineNumbs(in *bufio.Reader) []string {
	line, _ := in.ReadString('\n')
	line = strings.ReplaceAll(line, "\r", "")
	line = strings.ReplaceAll(line, "\n", "")
	numbs := strings.Split(line, " ")
	return numbs
}

func ReadArrInt(in *bufio.Reader) []int {
	numbs := ReadLineNumbs(in)
	arr := make([]int, len(numbs))
	for i, n := range numbs {
		val, _ := strconv.Atoi(n)
		arr[i] = val
	}
	return arr
}

func ReadArrInt64(in *bufio.Reader) []int64 {
	numbs := ReadLineNumbs(in)
	arr := make([]int64, len(numbs))
	for i, n := range numbs {
		val, _ := strconv.ParseInt(n, 10, 64)
		arr[i] = val
	}
	return arr
}
