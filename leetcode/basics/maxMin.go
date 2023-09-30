package basic

import (
	"fmt"
	"math"
)

func MinIntV(args ...int) int {
	fmt.Println(math.MaxInt)
	min := args[0]

	for _, v := range args {
		if min > v {
			min = v
		}
	}
	return min
}

func MinInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxIntV(args ...int) int {
	max := args[0]

	for _, v := range args {
		if max < v {
			max = v
		}
	}
	return max
}

func MaxInt(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func MaxFloat(args ...float64) float64 {
	max := args[0]

	for _, v := range args {
		if max < v {
			max = v
		}
	}

	return max
}

func MinFloat(args ...float64) float64 {
	min := args[0]

	for _, v := range args {
		if min > v {
			min = v
		}
	}

	return min
}
