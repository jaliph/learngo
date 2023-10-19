// https://leetcode.com/problems/find-two-non-overlapping-sub-arrays-each-with-target-sum/description/

package prefixsum

import "fmt"

func minSumOfLengths(arr []int, target int) int {

	prefixMap := map[int]int{}
	var length int = 1e9
	var ans int = 1e9

	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	prefixMap[0] = -1
	sum := 0
	for i, n := range arr {
		sum += n
		prefixMap[sum] = i
	}

	sum = 0
	for i, n := range arr {
		sum += n
		if v1, ok := prefixMap[sum-target]; ok {
			fmt.Println("First part", i, v1)
			length = Min(length, i-v1)
		}

		if v2, ok := prefixMap[sum+target]; ok && length < 1e9 {
			fmt.Println("Second part", v2, i)
			ans = Min(ans, length+(v2-i))
		}
	}

	return ans
}

func Driver() {
	arr := []int{3, 2, 2, 4, 3}
	target := 3
	fmt.Println(minSumOfLengths(arr, target))

}
