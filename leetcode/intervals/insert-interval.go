package intervals

// https://leetcode.com/problems/insert-interval/
func Insert(intervals [][]int, newInterval []int) [][]int {

	i := 0
	n := len(intervals)
	res := [][]int{}
	for i < n && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}

	Max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	Min := func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}

	for i < n && intervals[i][0] <= newInterval[1] {
		newInterval[0] = Min(newInterval[0], intervals[i][0])
		newInterval[1] = Max(newInterval[1], intervals[i][1])
		i++
	}
	res = append(res, newInterval)

	for i < n {
		res = append(res, intervals[i])
		i++
	}

	return res
}
