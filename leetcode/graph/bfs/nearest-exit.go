package graph

// https://leetcode.com/problems/nearest-exit-from-entrance-in-maze/

type N_Cell struct {
	x    int
	y    int
	dist int
}

func NewN_Cell(x, y, dist int) N_Cell {
	return N_Cell{x, y, dist}
}

func nearestExit(maze [][]byte, entrance []int) int {

	R := len(maze)
	C := len(maze[0])

	visited := make([][]bool, R)
	rows := make([]bool, R*C)
	for i := range visited {
		visited[i] = rows[i*C : (i+1)*C]
	}

	visited[entrance[0]][entrance[1]] = true
	q := []N_Cell{NewN_Cell(entrance[0], entrance[1], 0)}
	paths := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if curr.dist != 0 && (curr.x == 0 || curr.y == 0 || curr.x == R-1 || curr.y == C-1) {
			return curr.dist
		}

		for _, p := range paths {
			n_x := curr.x + p[0]
			n_y := curr.y + p[1]

			if n_x < 0 || n_y < 0 || n_x >= R || n_y >= C || visited[n_x][n_y] {
				continue
			}

			if maze[n_x][n_y] == '+' {
				continue
			}

			visited[n_x][n_y] = true
			q = append(q, NewN_Cell(n_x, n_y, curr.dist+1))
		}
	}

	return -1
}

// func Driver() {
// 	maze := [][]byte{{'+', '+', '+'}, {'.', '.', '.'}, {'+', '+', '+'}}
// 	entrance := []int{1, 0}
// 	fmt.Println(nearestExit(maze, entrance))
// }
