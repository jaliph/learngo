// https://leetcode.com/problems/constrained-subsequence-sum/
package stacks

func ConstrainedSubsetSum(nums []int, k int) int {
	deque := [][2]int{}

	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	total := 0
	res := 0
	for i, v := range nums {
		if len(deque) > 0 {
			total = v + deque[len(deque)-1][1]
		} else {
			total = v
		}

		res = Max(res, total)

		for len(deque) > 0 && total >= deque[len(deque)-1][1] {
			deque = deque[:len(deque)-1]
		}

		if total > 0 {
			deque = append(deque, [2]int{i, total})
		}

		if len(deque) > 0 && deque[0][0] < i-k {
			deque = deque[1:]
		}

		// fmt.Print(deque)
	}

	return res
}

// func Driver() {
// 	fmt.Println(constrainedSubsetSum([]int{10, 2, -10, 5, 20}, 2))
// }
