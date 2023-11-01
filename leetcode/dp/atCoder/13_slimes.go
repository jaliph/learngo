package dp

import "fmt"

func Slime(slimes []int) int {

	N := len(slimes)
	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	const INF int = 1e12

	var solve func(int, int) int
	solve = func(i, j int) int {
		if i > j {
			return 0
		}

		if i == j {
			return slimes[i]
		}

		var temp int = 1e9
		for k := i; k < j; k++ {
			left := solve(i, k)
			right := solve(k+1, j)
			fmt.Println(left, right)
			temp = Min(temp, right+left+slimes[i-1]+slimes[k])
		}
		fmt.Println(i, j, temp)
		return temp
	}

	return solve(1, N-1)
}

func Driver() {
	mats := []int{10, 20, 30, 40}
	fmt.Println(Slime(mats))
}
