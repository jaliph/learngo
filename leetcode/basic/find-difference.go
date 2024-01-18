package basic

// https://leetcode.com/problems/find-indices-with-index-and-value-difference-i/
// https://leetcode.com/problems/find-indices-with-index-and-value-difference-ii/

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	minIdx := 0
	maxIdx := 0

	Abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	for i := indexDifference; i < len(nums); i++ {
		if nums[i-indexDifference] < nums[minIdx] {
			minIdx = i - indexDifference
		}

		if nums[i-indexDifference] > nums[maxIdx] {
			maxIdx = i - indexDifference
		}

		if Abs(nums[i]-nums[minIdx]) >= valueDifference {
			return []int{minIdx, i}
		}

		if Abs(nums[i]-nums[maxIdx]) >= valueDifference {
			return []int{maxIdx, i}
		}
	}
	return []int{-1, -1}
}
