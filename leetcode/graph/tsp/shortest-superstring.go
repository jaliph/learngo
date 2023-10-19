package graph

import "fmt"

// INVALID SOLUTION
// https://leetcode.com/problems/find-the-shortest-superstring/solutions/195487/python-bfs-solution-with-detailed-explanation-with-extra-chinese-explanation/
// TSP
func ShortestSuperstring(words []string) string {

	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	overLapFinderString := func(str1, str2 string) string {
		str := str1 + str2
		var i int
		for i = 0; i < len(str1); i++ {
			if str1[i:] == str2[0:Min(len(str1)-i, len(str2))] {
				str = str1 + str2[len(str1)-i:]
				break
			}
		}
		return str
	}

	overLapFinder := func(str1, str2 string) int {
		var res int = len(str2)
		for i := 0; i < len(str1); i++ {
			if str1[i:] == str2[0:Min(len(str1)-i, len(str2))] {
				res = len(str1) - i
				break
			}
		}
		return res
	}

	g := [][]int{}
	for i, w1 := range words {
		row := make([]int, len(words))
		for j, w2 := range words {
			if i == j {
				continue
			}
			row[j] = overLapFinder(w1, w2)
		}
		g = append(g, row)
	}

	paths := []string{}
	visited := map[int]bool{}
	var solve func(int)
	solve = func(curr int) {

		neigh := INVALID
		cost := -INVALID
		visited[curr] = true

		for i, v := range g[curr] {
			_, ok := visited[i]
			if !ok && v != 0 && v > cost {
				neigh = i
				cost = v
			}
		}

		if neigh == INVALID {
			return
		}

		paths = append(paths, overLapFinderString(words[curr], words[neigh]))
		solve(neigh)
	}

	solve(0)

	fmt.Println(paths)
	return ""
}

func Driver() {
	words := []string{"catg", "ctaagt", "gcta", "ttca", "atgcatc"}
	fmt.Println(ShortestSuperstring(words))
}

// [catgcta gctaagt ctaagttca ttcatgcatc]
// [catgctaagt ctaagtatgcatc atgcatcgcta gctattca]
// ttcatgcatc
// gctaagttcatgcatc
// gctattcactaagtatgcatc
// atgcatcgctaagtatgcatc

// unc ShortestSuperstring(words []string) string {

// 	Min := func(a int, b int) int {
// 		if a < b {
// 			return a
// 		}
// 		return b
// 	}

// 	overLapFinder := func(str1, str2 string) string {
// 		str := str1 + str2
// 		var i int
// 		for i = 0; i < len(str1); i++ {
// 			if str1[i:] == str2[0:Min(len(str1)-i, len(str2))] {
// 				str = str1 + str2[len(str1)-i:]
// 				break
// 			}
// 		}
// 		return str
// 	}

// 	g := [][]string{}
// 	for i, w1 := range words {
// 		row := make([]string, len(words))
// 		for j, w2 := range words {
// 			if i == j {
// 				continue
// 			}
// 			row[j] = overLapFinder(w1, w2)
// 		}
// 		g = append(g, row)
// 	}

// 	paths := []string{}
// 	visited := map[int]bool{}
// 	var solve func(int)
// 	solve = func(curr int) {

// 		neigh := INVALID
// 		cost := INVALID
// 		visited[curr] = true

// 		for i, v := range g[curr] {
// 			_, ok := visited[i]
// 			if !ok && len(v) != 0 && len(v) < cost {
// 				neigh = i
// 				cost = len(v)
// 			}
// 		}

// 		if neigh == INVALID {
// 			return
// 		}

// 		paths = append(paths, g[curr][neigh])
// 		solve(neigh)
// 	}

// 	solve(0)

// 	fmt.Println(paths)
// 	return ""
// }
