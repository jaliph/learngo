package dp

import "math"

// always break in 3s
/*


    Why 3s? By testing various numbers, it becomes evident that the number 3 plays a significant role in maximizing the product. For example:
        4 = 2 + 2 (product = 4)
        5 = 2 + 3 (product = 6)
        6 = 3 + 3 (product = 9)
        7 = 3 + 2 + 2 (product = 12)
        8 = 3 + 3 + 2 (product = 18)
        9 = 3 + 3 + 3 (product = 27)

    Why not 4s? Anytime we have a 4, we can always split it into two 2s, which will give a better or equal product.

    Handling Remainders: When n is divided by 3, the remainder can be 0, 1, or 2:
        Remainder 0: n is a perfect multiple of 3. We break n entirely into 3s.
        Remainder 1: It's better to take a 3 out and add a 4 (which is split as two 2s). This is because the product of three numbers, 3, 1, and n-4, is less than the product of the two numbers, 4 and n-4.
        Remainder 2: We simply multiply the leftover 2 with the product of all 3s.

Complexity

    Time complexity: O(log n). This is because our primary operation involves raising 3 to the power of count_of_3s which, in Python, takes O(log n) time.

    Space complexity: O(1). We only use a constant amount of extra space regardless of the input size.

*/

func IntegerBreakMaths(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	count_of_3s := n / 3
	remainder := n % 3

	if remainder == 0 {
		return int(math.Pow(3, float64(count_of_3s)))
	} else if remainder == 1 {
		return int(math.Pow(3, float64(count_of_3s-1))) * 4
	} else {
		return int(math.Pow(3, float64(count_of_3s))) * 2
	}
}

func IntegerBreakDP(n int) int {
	dp := make(map[int]int)

	dp[1] = 1

	for i := 2; i <= n; i++ {
		if i == n {
			dp[i] = 0
		} else {
			dp[i] = i
		}
		for j := 1; j < i; j++ {
			dp[i] = Max(dp[i], dp[j]*dp[i-j])
		}
	}
	return dp[n]
}

func IntegerBreakMemoise(n int) int {
	var recur func(num int) int
	memoise := make(map[int]int)
	memoise[1] = 1
	recur = func(num int) int {
		if v, ok := memoise[num]; ok {
			return v
		}

		var res int

		if n == num {
			res = 0
		} else {
			res = num
		}
		for i := 1; i < num; i++ {
			val := recur(i) * recur(num-i)
			res = Max(res, val)
		}
		memoise[num] = res
		return res
	}

	return recur(n)
}

func IntegerBreakBrute(n int) int {
	var recur func(num int) int
	recur = func(num int) int {
		if num == 1 {
			return num
		}

		var res int
		// dont't go further after 1 break
		if n == num {
			res = 0
		} else {
			res = num
		}
		for i := 1; i < num; i++ {
			val := recur(i) * recur(num-i)
			res = Max(res, val)
		}
		return res
	}

	return recur(n)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
