package graph

import "fmt"

type BusNode struct {
	stop string
	cost int
}

func NewBusNode(bus string, cost int) BusNode {
	return BusNode{bus, cost}
}

func GetString(i int) string {
	return fmt.Sprintf("%d", i)
}

func numBusesToDestination(routes [][]int, source int, target int) int {

	g := map[string][]BusNode{}

	for bus, route := range routes {
		for _, stop := range route {
			if _, ok := g[GetString(stop)]; !ok {
				g[GetString(stop)] = []BusNode{}
			}
			if _, ok := g["#"+GetString(bus)]; !ok {
				g["#"+GetString(bus)] = []BusNode{}
			}
			g[GetString(stop)] = append(g[GetString(stop)], NewBusNode("#"+GetString(bus), 1))
			g["#"+GetString(bus)] = append(g["#"+GetString(bus)], NewBusNode(GetString(stop), 0))
		}
	}

	// fmt.Println(g)

	q := []BusNode{}
	visited := map[string]bool{}
	q = append(q, NewBusNode(GetString(source), 0))
	visited[GetString(source)] = true

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if curr.stop == GetString(target) {
			return curr.cost
		}

		for _, neigh := range g[curr.stop] {
			if _, ok := visited[neigh.stop]; ok {
				continue
			}
			visited[neigh.stop] = true
			q = append(q, NewBusNode(neigh.stop, neigh.cost+curr.cost))
		}
	}

	return -1
}

func Driver() {
	fmt.Println(numBusesToDestination([][]int{{1, 2, 7}, {3, 6, 7}}, 1, 6))
}
