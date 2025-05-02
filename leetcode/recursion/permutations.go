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

// Recur1
func Permute1(nums []int) [][]int {

	insertAti := func(arr *[]int, idx int, data int) {
		if idx == len(*arr) {
			*arr = append(*arr, data)
		} else {
			*arr = append((*arr)[:idx+1], (*arr)[idx:]...)
			(*arr)[idx] = data
		}
	}

	var recur func(perms []int) [][]int
	recur = func(perms []int) [][]int {
		if len(perms) == 0 {
			return [][]int{{}}
		}

		n := perms[0]
		sub_perms := recur(perms[1:])

		new_perms := [][]int{}
		for _, p := range sub_perms {
			for idx := 0; idx <= len(p); idx++ {
				copy_perm := make([]int, len(p))
				copy(copy_perm, p)
				insertAti(&copy_perm, idx, n)
				new_perms = append(new_perms, copy_perm)
			}
		}
		return new_perms
	}

	return recur(nums)
}

func Permute2(nums []int) [][]int {

	insertAti := func(arr *[]int, idx int, data int) {
		if idx == len(*arr) {
			*arr = append(*arr, data)
		} else {
			*arr = append((*arr)[:idx+1], (*arr)[idx:]...)
			(*arr)[idx] = data
		}
	}

	perms := [][]int{{}}

	for _, n := range nums {
		new_perms := [][]int{}
		for _, p := range perms {
			for idx := 0; idx <= len(p); idx++ {
				copy_perm := make([]int, len(p))
				copy(copy_perm, p)
				insertAti(&copy_perm, idx, n)
				new_perms = append(new_perms, copy_perm)
			}
		}
		perms = new_perms
	}
	return perms
}
