package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var nrchars, nrwords, nrlines int
var InputReader *bufio.Reader

func main() {
	InputReader = bufio.NewReader(os.Stdin)

	str, err := InputReader.ReadString('S')
	if err != nil {
		panic(err)
	}
	Counters(str)
	fmt.Println(nrchars, nrwords, nrlines)
}

func Counters(input string) {
	c := []byte(input)
	nrchars = len(c)
	nrwords = len(strings.Split(input, " "))
	nrlines = len(strings.Split(input, "\n"))
}
