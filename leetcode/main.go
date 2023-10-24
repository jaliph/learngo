package main

import (
	"fmt"
)

func main() {
	if float64(0.1)+float64(0.2) == float64(0.3) {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}
}
