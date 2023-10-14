// https://leetcode.com/problems/maximum-number-of-events-that-can-be-attended/description/

package intervals

import (
	"fmt"
	"sort"
)

func maxEvents(events [][]int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	count := 0

	totalDays := func(e [][]int) int {
		max := 0
		for _, v := range e {
			if max < v[1] {
				max = v[1]
			}
		}
		return max
	}(events)

	h := NewHeap(func(i, j int) int {
		return i - j
	})

	idx := 0
	for d := 1; d <= totalDays; d++ {
		for idx < len(events) && events[idx][0] <= d {
			h.Push(events[idx][1])
			idx++
		}

		for h.size > 0 && h.Peek() < d {
			h.Pop()
		}

		if h.size > 0 {
			h.Pop()
			count++
		} else if idx >= len(events) {
			break
		}
	}

	return count
}

func Driver() {
	events := [][]int{{1, 2}, {2, 3}, {3, 4}}

	fmt.Println(maxEvents(events))
}
