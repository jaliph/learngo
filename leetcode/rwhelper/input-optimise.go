// https://codeforces.com/blog/entry/94245

// go build -o main
// ./main -cpuprofile=out.prof < input_large
// go tool pprof main out.prof
package rwhelper

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
)

func Input_Optimised() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	in := bufio.NewReader(os.Stdin)
	tc := Readint(in)

	for i := 0; i < tc; i++ {
		n := Readint(in)
		_ = ReadArrInt(in)
		fmt.Println("N: ", n)
		fmt.Println("DONE TC: ", i+1)
	}
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
