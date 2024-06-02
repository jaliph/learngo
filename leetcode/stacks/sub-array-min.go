package stacks

// https://leetcode.com/problems/sum-of-subarray-minimums/

func sumSubarrayMinsBrute(arr []int) int {

	sum := 0

	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := 0; i < len(arr); i++ {
		min := arr[i]
		for j := i; j < len(arr); i++ {
			min = Min(min, arr[j])
			sum += min
		}
	}
	return sum
}

func sumSubarrayMins(arr []int) int {
	// this forces everything out of the stack at the end
	arr = append(arr, -1)
	min := NewStack()
	sum := 0
	for i := range arr {
		for min.Len() > 0 && arr[i] < arr[min.Top()] {
			//[p.....][c][......i]
			//    P           K
			// the number of subarrays with c as the minimum is P * K i.e. (c-p) * (i-c) see bottom
			c, p := min.Pop(), min.Top()
			P, K := c-p, i-c
			sum += arr[c] * P * K // see proof below
		}

		min.Push(i)
	}

	return sum % int(1e9+7)
}

type stack struct {
	Push func(x int)
	Pop  func() int
	Top  func() int
	Len  func() int
}

func NewStack() stack {
	arr := make([]int, 0)
	return stack{
		Push: func(x int) {
			arr = append(arr, x)
		},
		Pop: func() int {
			item := arr[len(arr)-1]
			arr = arr[:len(arr)-1]
			return item
		},
		Top: func() int {
			if len(arr) == 0 {
				return -1
			}
			return arr[len(arr)-1]
		},
		Len: func() int {
			return len(arr)
		},
	}
}

func sumSubarrayMins2(arr []int) int {
	const MOD int = 1e9 + 7
	n := len(arr)
	left := make([]int, n)
	right := make([]int, n)

	for i := range left {
		left[i] = -1
		right[i] = n
	}

	stack := []int{}

	for i := 0; i < n; i++ {
		for len(stack) > 0 && arr[stack[len(stack)-1]] >= arr[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	stack = []int{}
	for i := len(arr) - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[stack[len(stack)-1]] > arr[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	// fmt.Println(left)
	// fmt.Println(right)

	sum := 0
	for i := 0; i < n; i++ {
		sum += ((i - left[i]) * (right[i] - i) * arr[i]) % MOD
		sum %= MOD
	}

	return sum
}

// func Driver() {
// 	arr := []int{3, 1, 2, 4}
// 	fmt.Println(sumSubarrayMins(arr))
// }

// https://leetcode.com/problems/sum-of-subarray-minimums/solutions/2846397/monostack/
