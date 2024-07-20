package stacks

type Entry struct {
	idx int
	val int
}

func MaxSumMinProduct(nums []int) int {
	const MOD = 1e9 + 7
	prefix := make([]int, len(nums)+1)

	for i, n := range nums {
		prefix[i+1] = prefix[i] + n
	}
	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	res := 0
	stack := []Entry{}

	for i, v := range nums {
		newStart := i
		for len(stack) > 0 && stack[len(stack)-1].val > v {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			sum := prefix[i] - prefix[top.idx]
			res = Max(res, (sum * top.val))
			newStart = top.idx
		}
		stack = append(stack, Entry{val: v, idx: newStart})
	}

	for _, entry := range stack {
		sum := prefix[len(nums)] - prefix[entry.idx]
		res = Max(res, (sum * entry.val))
	}

	return res % MOD
}
