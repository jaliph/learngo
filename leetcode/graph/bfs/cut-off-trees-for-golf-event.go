package graph

import (
	"sort"
)

// https://leetcode.com/problems/cut-off-trees-for-golf-event/

type Coord struct {
	x   int
	y   int
	val int
}

func CutOffTree(forest [][]int) int {

	R := len(forest)
	C := len(forest[0])

	// tMap := map[Coord]int{}
	trees := []Coord{}

	for i := range forest {
		for j := range forest[i] {
			if forest[i][j] > 1 {
				trees = append(trees, Coord{i, j, forest[i][j]})
			}
		}
	}

	sort.Slice(trees, func(i, j int) bool {
		return trees[i].val < trees[j].val
	})

	paths := [][]int{
		{1, 0},
		{0, 1},
		{0, -1},
		{-1, 0},
	}

	bfs := func(sx, sy, dx, dy int) int {
		row := R
		col := C
		visited := make([][]bool, row)
		rows := make([]bool, row*col)
		for i := range visited {
			visited[i] = rows[i*C : (i+1)*C]
		}

		q := [][3]int{}
		q = append(q, [3]int{sx, sy, 0})
		visited[sx][sy] = true

		for len(q) > 0 {
			curr := q[0]
			q = q[1:]

			if curr[0] == dx && curr[1] == dy {
				return curr[2]
			}

			for _, p := range paths {
				nx := p[0] + curr[0]
				ny := p[1] + curr[1]

				if nx >= 0 && nx < R && ny >= 0 && ny < C && !visited[nx][ny] && forest[nx][ny] != 0 {
					visited[nx][ny] = true
					q = append(q, [3]int{nx, ny, curr[2] + 1})
				}
			}
		}
		return -1
	}

	count := 0
	start := Coord{x: 0, y: 0}
	for i := 0; i < len(trees); i++ {
		dist := bfs(start.x, start.y, trees[i].x, trees[i].y)
		if dist == -1 {
			return -1
		}
		start.x = trees[i].x
		start.y = trees[i].y
		count += dist
	}
	return count
}
