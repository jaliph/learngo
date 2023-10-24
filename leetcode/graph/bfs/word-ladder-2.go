// https://leetcode.com/problems/word-ladder-ii/
package graph

import (
	"math"
)

func findLadders(beginWord string, endWord string, wordList []string) [][]string {

	// check endword exists int he wordlist
	idx := len(wordList)
	for v, w := range wordList {
		if endWord == w {
			idx = v
			break
		}
	}
	if idx == len(wordList) {
		return [][]string{}
	}

	// prepare the graph
	g := make(map[string][]string)

	for _, w := range wordList {
		linkGen(w, g)
	}

	// prepare the for bfs
	dist := make(map[string]int)
	parents := make(map[string][]string)
	q := []string{}

	q = append(q, beginWord)
	dist[beginWord] = 0
	parents[beginWord] = []string{"-1"}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		for i := 0; i < len(curr); i++ {
			link := string(curr[0:i]) + "*" + string(curr[i+1:])
			for _, v := range g[link] {
				// set the default values for v
				if _, ok := dist[v]; !ok {
					dist[v] = math.MaxInt
				}

				// main logic
				if dist[v] > dist[curr]+1 {
					dist[v] = dist[curr] + 1
					parents[v] = []string{curr}
					q = append(q, v) // push the child
				} else if dist[v] == dist[curr]+1 {
					parents[v] = append(parents[v], curr)
				}
			}
		}
	}

	// find all the combinations
	res := [][]string{}
	paths := []string{}
	var com func(node string)
	com = func(node string) {
		// base case
		if node == "-1" {
			result := make([]string, len(paths)-1) // removing the last "-1"
			copy(result, paths[:len(paths)-1])
			reverse(result)
			res = append(res, result)
			return
		}
		// recur
		for _, par := range parents[node] {
			paths = append(paths, par)
			com(par)
			paths = paths[:len(paths)-1]
		}
	}
	paths = append(paths, endWord)
	com(endWord)

	return res
}

func linkGen(w string, g map[string][]string) {
	for i := 0; i < len(w); i++ {
		link := string(w[0:i]) + "*" + string(w[i+1:])
		if _, ok := g[link]; !ok {
			g[link] = []string{}
		}
		g[link] = append(g[link], w)
	}
}

func reverse(strs []string) {
	l, r := 0, len(strs)-1
	for l < r {
		strs[l], strs[r] = strs[r], strs[l]
		l++
		r--
	}
}
