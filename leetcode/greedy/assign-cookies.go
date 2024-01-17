// https://leetcode.com/problems/assign-cookies/

package greedy

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	g_ptr := len(g) - 1
	s_ptr := len(s) - 1

	cnt := 0
	for g_ptr > 0 && s_ptr > 0 {
		if s[s_ptr] >= g[g_ptr] {
			cnt++
			s_ptr--
			g_ptr--
		} else {
			g_ptr--
		}
	}

	return cnt
}
