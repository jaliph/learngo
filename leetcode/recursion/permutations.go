package recursion

import "fmt"

// https://leetcode.com/problems/permutations/

func Permute(nums []int) [][]int {

	insertAti := func(arr []int, idx int, data int) []int {
		if idx == len(arr) {
			arr = append(arr, data)
		} else {
			arr = append(arr[:idx+1], arr[idx:]...)
			arr[idx] = data
		}
		return arr
	}

	var dfs func(nums []int) [][]int
	dfs = func(nums []int) [][]int {
		fmt.Println(nums)
		if len(nums) == 0 {
			return [][]int{{}}
		}

		first := nums[0]
		res := dfs(nums[1:])

		copyArr := [][]int{}
		for _, arr := range res {
			for i := 0; i <= len(arr); i++ {
				newArr := make([]int, len(arr))
				copy(newArr, arr)
				newArr = insertAti(newArr, i, first)
				copyArr = append(copyArr, newArr)
			}
		}
		return copyArr
	}
	return dfs(nums)
}
