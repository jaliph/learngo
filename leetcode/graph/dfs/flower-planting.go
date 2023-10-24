package graph

// https://leetcode.com/problems/flower-planting-with-no-adjacent/
import (
	"fmt"
)

// m-coloring
func GardenNoAdj(n int, paths [][]int) []int {

	// setup graph
	g := make([][]int, n)

	for _, edges := range paths {
		if len(g[edges[0]-1]) == 0 {
			g[edges[0]-1] = []int{}
		}
		if len(g[edges[1]-1]) == 0 {
			g[edges[1]-1] = []int{}
		}
		g[edges[0]-1] = append(g[edges[0]-1], edges[1]-1)
		g[edges[1]-1] = append(g[edges[1]-1], edges[0]-1)
	}
	// fmt.Println(g)

	colored := make([]int, n)
	// return []int{0}

	isSafe := func(vertex, color int) bool {
		for _, v := range g[vertex] {
			if colored[v] == color {
				return false
			}
		}
		fmt.Println("safe", vertex, color)
		return true
	}

	var mcoloring func(int) bool
	mcoloring = func(vtx int) bool {
		fmt.Println("visiting ", vtx)
		// base
		if vtx == n {
			return true
		}

		// recur

		// check for each color
		for color := 1; color <= 4; color++ {
			if isSafe(vtx, color) {
				colored[vtx] = color
				if mcoloring(vtx + 1) {
					return true
				}
				colored[vtx] = 0
			}
		}
		return false
	}

	mcoloring(0)
	return colored
}

// func Driver() {
// 	n := 3
// 	paths := [][]int{{1, 2}, {2, 3}, {3, 1}}
// 	fmt.Println(gardenNoAdj(n, paths))
// }
