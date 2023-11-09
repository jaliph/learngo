package stacks

// https://leetcode.com/problems/find-the-winner-of-an-array-game/

func getWinnertWinner(arr []int, k int) int {

	MaxV := func(args ...int) int {
		max := args[0]

		for i := 1; i < len(args); i++ {
			if max < args[i] {
				max = args[i]
			}
		}
		return max
	}

	if k >= len(arr) {
		return MaxV(arr...)
	}

	p1 := arr[0]
	arr = arr[1:]
	x := k

	for k > 0 {
		p2 := arr[0]
		arr = arr[1:]
		if p1 > p2 {
			k--
			arr = append(arr, p2)
		} else {
			k = x
			k--
			arr = append(arr, p1)
			p1 = p2
		}
	}
	return p1
}

// func Driver() {
// 	fmt.Println(getWinner([]int{2, 1, 3, 5, 4, 6, 7}, 2))
// }
