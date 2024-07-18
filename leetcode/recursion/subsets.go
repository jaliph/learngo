package recursion

func Subsets(nums []int) [][]int {
	res := [][]int{{}}

	for _, n := range nums {
		currLength := len(res)
		for i := 0; i < currLength; i++ {
			cp := make([]int, len(res[i]))
			copy(cp, res[i])
			cp = append(cp, n)
			res = append(res, cp)
		}
	}

	return res
}
