package rwhelper

import (
	"fmt"
	"os"
)

var f *os.File

const FILE = "out.txt"

func Write(v ...interface{}) {
	if f == nil {
		os.Truncate(FILE, 0)
		f, _ = os.OpenFile(FILE, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	}

	fmt.Fprintln(f, v...)
}

func CloseWriter() {
	f.Close()
}
