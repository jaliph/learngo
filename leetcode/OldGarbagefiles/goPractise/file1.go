package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Println(sum(2, 3))
}

func sum(a int, b int) int {
	return a + b
}
