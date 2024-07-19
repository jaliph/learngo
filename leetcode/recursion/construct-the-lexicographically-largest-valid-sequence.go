package recursion

// https://leetcode.com/problems/construct-the-lexicographically-largest-valid-sequence/
func dfs(nums []int, arr []int, p int) []int {
	if p == len(arr) {
		return arr
	}
	if arr[p] != 0 {
		return dfs(nums, arr, p+1)
	}
	for i := len(nums) - 1; i >= 1; i-- {
		dist := i
		if i == 1 {
			dist = 0
		}
		if nums[i] != 0 && dist+p < len(arr) && arr[dist+p] == 0 {
			arr[p] = i
			arr[p+dist] = i
			nums[i] = 0
			res := dfs(nums, arr, p+1)
			if len(res) > 0 {
				return res
			}
			nums[i] = 1
			arr[p] = 0
			arr[p+dist] = 0
		}
	}
	return []int{}
}

func ConstructDistancedSequence(n int) []int {
	arr := make([]int, (n-1)*2+1)
	nums := make([]int, n+1)
	for i := range nums {
		nums[i] = 1
	}
	return dfs(nums, arr, 0)
}
