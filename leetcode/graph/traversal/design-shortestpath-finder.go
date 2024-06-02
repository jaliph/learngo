package graph

const INF int = 1e12

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

type Graph struct {
	g [][]int
	n int
}

func Constructor(n int, edges [][]int) Graph {
	R := n
	C := n
	g := make([][]int, R)
	rows := make([]int, R*C)

	for i := range rows {
		rows[i] = INF
	}
	for i := range g {
		g[i] = rows[i*C : (i+1)*C]
	}

	for i := range g {
		g[i][i] = 0
	}

	for _, edge := range edges {
		g[edge[0]][edge[1]] = edge[2]
	}

	// RUN Flyod Warshals
	for k := 0; k < n; k++ {

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {

				if i == j {
					continue
				}

				g[i][j] = Min(g[i][j], g[i][k]+g[k][j])
			}
		}
	}

	return Graph{g, n}
}

func (this *Graph) AddEdge(edge []int) {
	if this.g[edge[0]][edge[1]] <= edge[2] {
		return
	}

	this.g[edge[0]][edge[1]] = edge[2]

	// RUN Flyod Warshals
	for i := 0; i < this.n; i++ {
		for j := 0; j < this.n; j++ {

			this.g[i][j] = Min(this.g[i][j], this.g[i][edge[0]]+edge[2]+this.g[edge[1]][j])
		}
	}
}

func (this *Graph) ShortestPath(node1 int, node2 int) int {
	if this.g[node1][node2] == INF {
		return -1
	}
	return this.g[node1][node2]
}

// func Driver() {
// 	// g := Constructor(4, [][]int{{0, 2, 5}, {0, 1, 2}, {1, 2, 1}, {3, 0, 3}})

// 	g := Constructor(17, [][]int{{6, 3, 475731}, {3, 7, 515993}, {13, 8, 904914}, {1, 15, 665336}, {3, 4, 419955}, {2, 5, 591627}, {15, 13, 180374}, {8, 6, 939578}, {7, 10, 459913}, {8, 1, 942800}, {14, 6, 876558}, {15, 5, 215248}, {13, 14, 762427}, {1, 5, 914567}, {6, 12, 919273}, {12, 16, 342212}, {10, 8, 60822}, {3, 14, 947396}, {15, 0, 263172}, {10, 6, 711514}, {9, 14, 120577}, {11, 5, 476839}, {11, 13, 438668}, {12, 9, 527842}, {14, 16, 588402}, {15, 2, 195790}, {1, 9, 785659}, {2, 7, 787223}, {11, 7, 99316}, {6, 1, 948004}, {4, 12, 31559}, {5, 4, 453573}, {5, 8, 141131}, {5, 12, 767697}, {1, 12, 312956}, {14, 11, 374561}, {15, 11, 19433}, {0, 9, 227239}, {12, 10, 325409}, {16, 13, 413278}, {10, 1, 155788}, {5, 3, 294209}, {7, 5, 54490}, {3, 13, 701716}, {2, 8, 927178}, {12, 14, 367956}, {14, 10, 953761}, {3, 16, 9061}, {2, 3, 421223}, {4, 15, 189155}, {14, 2, 711954}, {16, 15, 734202}, {7, 4, 917325}, {0, 1, 818496}, {8, 3, 548282}, {16, 12, 385482}})

// 	fmt.Println(g.ShortestPath(0, 11))
// }
