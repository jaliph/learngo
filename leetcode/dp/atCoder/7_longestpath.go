package dp

// https://atcoder.jp/contests/dp/tasks/dp_g

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LongestPath() {
	// Input Helpers
	ReadLineNumbs := func(in *bufio.Reader) []string {
		line, _ := in.ReadString('\n')
		line = strings.ReplaceAll(line, "\r", "")
		line = strings.ReplaceAll(line, "\n", "")
		numbs := strings.Split(line, " ")
		return numbs
	}

	ReadArrInt := func(in *bufio.Reader) []int {
		numbs := ReadLineNumbs(in)
		arr := make([]int, len(numbs))
		for i, n := range numbs {
			val, _ := strconv.Atoi(n)
			arr[i] = val
		}
		return arr
	}

	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	in := bufio.NewReader(os.Stdin)

	line1 := ReadArrInt(in)
	n := line1[0]
	e := line1[1]
	g := make([][]int, n)
	indegree := make([]int, n)

	for i := 0; i < e; i++ {
		edge := ReadArrInt(in)
		// fmt.Println(edge)
		if len(g[edge[0]-1]) == 0 {
			g[edge[0]-1] = []int{}
		}

		indegree[edge[1]-1]++
		g[edge[0]-1] = append(g[edge[0]-1], edge[1]-1)
	}

	longestPath := 0
	memoise := map[int]int{}

	var solve func(int) int
	solve = func(curr int) int {
		dist := 0

		if v, ok := memoise[curr]; ok {
			return v
		}

		for _, neigh := range g[curr] {
			dist = Max(dist, 1+solve(neigh))
		}
		memoise[curr] = dist
		return dist
	}

	for i, v := range indegree {
		if v == 0 {
			longestPath = Max(longestPath, solve(i))
		}
	}

	fmt.Println(longestPath)
}

/// Attempt 1

/**

type SpNode struct {
	v    int
	mask int
	dist int
}

func NewSpNode(vertex, mask, dist int) SpNode {
	return SpNode{vertex, mask, dist}
}

func LongestPath() {
	// Input Helpers
	ReadLineNumbs := func(in *bufio.Reader) []string {
		line, _ := in.ReadString('\n')
		line = strings.ReplaceAll(line, "\r", "")
		line = strings.ReplaceAll(line, "\n", "")
		numbs := strings.Split(line, " ")
		return numbs
	}

	ReadArrInt := func(in *bufio.Reader) []int {
		numbs := ReadLineNumbs(in)
		arr := make([]int, len(numbs))
		for i, n := range numbs {
			val, _ := strconv.Atoi(n)
			arr[i] = val
		}
		return arr
	}

	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	in := bufio.NewReader(os.Stdin)

	line1 := ReadArrInt(in)
	n := line1[0]
	e := line1[1]
	g := make([][]int, n)
	indegree := make([]int, n)
	visited := map[[2]int]bool{}

	for i := 0; i < e; i++ {
		edge := ReadArrInt(in)
		// fmt.Println(edge)
		if len(g[edge[0]-1]) == 0 {
			g[edge[0]-1] = []int{}
		}

		indegree[edge[1]-1]++
		g[edge[0]-1] = append(g[edge[0]-1], edge[1]-1)
	}
	q := []SpNode{}

	for i, v := range indegree {
		if v == 0 {
			q = append(q, NewSpNode(i, 1<<i, 0))
			visited[[2]int{i, 1 << i}] = true
		}
	}

	longestPath := 0

	for len(q) > 0 {
		curr := q[0]
		fmt.Println(curr)
		q = q[1:]
		longestPath = Max(longestPath, curr.dist)

		for _, neigh := range g[curr.v] {
			newMask := curr.mask | (1 << neigh)
			if _, ok := visited[[2]int{neigh, newMask}]; !ok {
				visited[[2]int{neigh, newMask}] = true
				q = append(q, NewSpNode(neigh, newMask, curr.dist+1))
			}
		}
	}

	fmt.Println(longestPath)

}

**/
