package rwhelper

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func Input_Brute() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	var tc int
	fmt.Scan(&tc)

	for i := 0; i < tc; i++ {
		var n int
		fmt.Scan(&n)
		arr := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&arr[j])
		}
		fmt.Println("DONE TC: ", i+1)
	}
}
