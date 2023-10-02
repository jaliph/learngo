package rwhelper

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var in *bufio.Reader

func NewInputReader() {
	if in == nil {
		in = bufio.NewReader(os.Stdin)
	}
}

func ReadString() string {
	NewInputReader()
	nStr, _ := in.ReadString('\n')
	nStr = strings.ReplaceAll(nStr, "\r", "")
	nStr = strings.ReplaceAll(nStr, "\n", "")
	return nStr
}

func ReadInt() int {
	NewInputReader()
	nStr, _ := in.ReadString('\n')
	nStr = strings.ReplaceAll(nStr, "\r", "")
	nStr = strings.ReplaceAll(nStr, "\n", "")
	n, _ := strconv.Atoi(nStr)
	return n
}

func ReadLineArray() []string {
	NewInputReader()
	line, _ := in.ReadString('\n')
	line = strings.ReplaceAll(line, "\r", "")
	line = strings.ReplaceAll(line, "\n", "")
	numbs := strings.Split(line, " ")
	return numbs
}

func ReadArrayInt() []int {
	NewInputReader()
	numbs := ReadLineNumbs(in)
	arr := make([]int, len(numbs))
	for i, n := range numbs {
		val, _ := strconv.Atoi(n)
		arr[i] = val
	}
	return arr
}

func ReadArrayInt64() []int64 {
	NewInputReader()
	numbs := ReadLineNumbs(in)
	arr := make([]int64, len(numbs))
	for i, n := range numbs {
		val, _ := strconv.ParseInt(n, 10, 64)
		arr[i] = val
	}
	return arr
}

func ReadArrayFloat64() []float64 {
	NewInputReader()
	numbs := ReadLineNumbs(in)
	arr := make([]float64, len(numbs))
	for i, n := range numbs {
		val, _ := strconv.ParseFloat(n, 64)
		arr[i] = val
	}
	return arr
}
