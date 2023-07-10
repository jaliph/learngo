package main

import "fmt"
import pk "book/pack1"

func main() {
	fmt.Println(fibarray(10))
}

func fibarray(term int) []int {
	fib := make([]int, term)
	fib[0] = 0
	fib[1] = 1
	for i := 2; i < term; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib
}


func enlarge(s []int, factor int) []int {
	s := pk.
	return s
}
