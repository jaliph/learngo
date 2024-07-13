package dp

// https://leetcode.com/problems/stickers-to-spell-word/
func MinStickers(stickers []string, target string) int {
	const INF int = 1e9
	sMaps := make([]map[string]int, len(stickers))

	mapMaker := func(str string) map[string]int {
		m := map[string]int{}
		for _, ch := range str {
			m[string(ch)]++
		}
		return m
	}

	copyMap := func(orig map[string]int) map[string]int {
		copy := map[string]int{}
		for k, v := range orig {
			copy[k] = v
		}
		return copy
	}

	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i, sticker := range stickers {
		sMaps[i] = mapMaker(sticker)
	}

	memoise := map[string]int{}
	var minFinder func(str string, m map[string]int) int

	minFinder = func(t string, m map[string]int) int {

		if val, ok := memoise[t]; ok {
			return val
		}

		var res int
		if len(m) > 0 {
			res = 1
		} else {
			res = 0
		}

		remaining := ""
		for _, ch := range t {
			if val, ok := m[string(ch)]; ok && val > 0 {
				m[string(ch)] = m[string(ch)] - 1
			} else {
				remaining += string(ch)
			}
		}

		if len(remaining) > 0 {
			// fmt.Println(remaining)
			used := INF
			for _, sMap := range sMaps {
				if _, ok := sMap[string(remaining[0])]; ok {
					used = Min(used, minFinder(remaining, copyMap(sMap)))
				}
			}
			memoise[remaining] = used
			res += used
		}

		return res
	}

	count := minFinder(target, map[string]int{})

	// fmt.Println(sMaps)
	if count == INF {
		return -1
	}
	return count
}
