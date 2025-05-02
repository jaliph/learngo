package basic

func LongestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	Max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	s := map[int]bool{}

	for _, n := range nums {
		s[n] = true
	}

	res := 1
	var temp int
	var count int
	for _, n := range nums {
		if s[n-1] {
			continue
		} else {
			count = 1
			temp = n + 1
			for s[temp] {
				temp += 1
				count += 1
			}
			res = Max(res, count)
			if res > len(nums)/2 {
				break
			}
		}
	}
	return res
}
