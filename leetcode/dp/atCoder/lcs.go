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

	str1 := ReadString(in)
	str2 := ReadString(in)

	fmt.Println(str1, str2)
}
