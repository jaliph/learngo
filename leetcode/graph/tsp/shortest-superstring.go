package graph

import "fmt"

// INVALID SOLUTION
// https://leetcode.com/problems/find-the-shortest-superstring/solutions/195487/python-bfs-solution-with-detailed-explanation-with-extra-chinese-explanation/
// TSP
type Node_SSS struct {
	node    int
	mask    int
	paths   []int
	overlap int
}

func NewNode_SSS(node, mask int, p []int, overlap int) Node_SSS {
	return Node_SSS{node, mask, p, overlap}
}

func ShortestSuperstring(words []string) string {
	n := len(words)
	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	overLapFinder := func(src, dst string) int {
		var res int = 1
		for i := 0; i < len(src); i++ {
			currentLen := len(src) - i
			if dst[0:Min(currentLen, len(dst))] == src[i:] {
				res = currentLen + 1
				break
			}
		}
		return res
	}

	g := make([][]int, n)
	rows := make([]int, n*n)
	for i := range g {
		g[i] = rows[i*n : (i+1)*n]
	}

	// Setup Graph
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			g[i][j] = overLapFinder(words[i], words[j])
			g[j][i] = overLapFinder(words[j], words[i])
		}
	}

	fmt.Println(g)
	// [[1, 2, 1, 2, 1], [2, 1, 1, 1, 1], [1, 1, 1, 1, 1], [1, 1, 1, 1, 1], [1, 1, 1, 1, 1]]

	R := 1 << n
	C := n

	dp := make([][]int, R)
	rows = make([]int, R*C)
	for i := range dp {
		dp[i] = rows[i*C : (i+1)*C]
	}

	// BFS = [node, mask, paths, overlap]
	q := []Node_SSS{}

	for i := 0; i < n; i++ {
		q = append(q, NewNode_SSS(i, 1<<i, []int{i}, 0))
	}

	// over all paths
	finallLen := -1
	finalPaths := []int{}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		// if overlap not good skip it
		if dp[curr.mask][curr.node] > curr.overlap {
			continue
		}

		if curr.mask == (1<<n)-1 && curr.overlap > finallLen {
			fmt.Println(curr)
			finallLen = curr.overlap
			finalPaths = curr.paths
			continue
		}

		for neigh := 0; neigh < n; neigh++ {
			if ((1 << neigh) & curr.mask) > 0 {
				// already visited
				continue
			}
			nextMask := curr.mask | (1 << neigh)
			nextoverlap := dp[curr.mask][curr.node] + g[curr.node][neigh]

			if nextoverlap > dp[nextMask][neigh] {
				dp[nextMask][neigh] = nextoverlap
				newPath := make([]int, len(curr.paths))
				copy(newPath, curr.paths)
				newPath = append(newPath, neigh)
				q = append(q, NewNode_SSS(neigh, nextMask, newPath, nextoverlap))
			}
		}
	}
	fmt.Println(finallLen)
	fmt.Println(finalPaths)

	ans := words[finalPaths[0]]
	for i := 1; i < len(finalPaths); i++ {
		prev, curr := finalPaths[i-1], finalPaths[i]
		overLap := g[prev][curr] - 1
		fmt.Println(overLap, words[curr])
		ans += words[curr][overLap:]
	}

	return ans
}

func Driver() {
	fmt.Println(ShortestSuperstring([]string{"ab", "a", "b"}))
}
