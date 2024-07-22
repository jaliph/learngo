package stacks

import "sort"

type P struct {
	pos int
	idx int
}

// https://leetcode.com/problems/robot-collisions/
func SurvivedRobotsHealths(positions []int, healths []int, directions string) []int {
	ps := []P{}
	for i, p := range positions {
		ps = append(ps, P{pos: p, idx: i})
	}

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].pos < ps[j].pos
	})

	stack := []int{}
	for _, v := range ps {
		i := v.idx
		if directions[i] == 'R' {
			stack = append(stack, i)
		} else {
			// -> -> <-
			// top is right, incoming is left..
			// loop should continue only left has +ve health healths[i] > 0

			for len(stack) > 0 && directions[stack[len(stack)-1]] == 'R' && healths[i] > 0 {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if healths[top] > healths[i] {
					healths[top] = healths[top] - 1
					healths[i] = 0
					// top should be put back for next iteration
					stack = append(stack, top)
				} else if healths[top] < healths[i] {
					healths[top] = 0
					healths[i] = healths[i] - 1
				} else {
					healths[top] = 0
					healths[i] = 0
				}
			}

			if healths[i] > 0 {
				stack = append(stack, i)
			}
		}
	}

	res := []int{}
	for _, h := range healths {
		if h > 0 {
			res = append(res, h)
		}
	}

	return res
}
