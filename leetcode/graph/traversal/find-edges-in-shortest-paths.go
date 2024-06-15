package graph

import (
	"container/heap"
)

type FANode struct {
	dist int
	v    int
}

type PQ []*FANode

func (pq *PQ) Len() int {
	return len(*pq)
}
func (pq *PQ) Less(i, j int) bool {
	return (*pq)[i].dist < (*pq)[j].dist
}
func (pq *PQ) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}
func (pq *PQ) Push(x interface{}) {
	node := x.(*FANode)
	*pq = append(*pq, &FANode{
		dist: node.dist,
		v:    node.v,
	})
}
func (pq *PQ) Pop() interface{} {
	old := *pq
	item := old[len(old)-1]
	*pq = old[:len(old)-1]
	return item
}

func FindAnswer(n int, edges [][]int) []bool {

	g := map[int][][3]int{}
	for i, e := range edges {
		g[e[0]] = append(g[e[0]], [3]int{e[1], e[2], i})
		g[e[1]] = append(g[e[1]], [3]int{e[0], e[2], i})
	}

	const INF int = 1e9 + 7
	dist := make([]int, n)
	for i := range dist {
		dist[i] = INF
	}
	dist[0] = 0
	q := PQ{}
	q.Push(&FANode{
		dist: 0,
		v:    0,
	})
	heap.Init(&q)
	for q.Len() > 0 {
		node := heap.Pop(&q).(*FANode)
		// Optimisation
		if dist[node.v] == node.dist {
			for _, n := range g[node.v] {
				neigh := n[0]
				wt := n[1]
				if wt+node.dist < dist[neigh] {
					dist[neigh] = wt + node.dist
					heap.Push(&q, &FANode{
						dist: dist[neigh],
						v:    neigh,
					})
				}
			}
		}
	}

	// fmt.Println(dist)

	heap.Push(&q, &FANode{
		dist: dist[n-1],
		v:    n - 1,
	})

	res := make([]bool, len(edges))
	visited := make([]bool, n)
	for q.Len() > 0 {
		node := heap.Pop(&q).(*FANode)
		for _, n := range g[node.v] {
			neig := n[0]
			wt := n[1]
			index := n[2]
			// Back Track to find the nodes
			if node.dist-wt == dist[neig] {
				if !visited[neig] {
					visited[neig] = true

					heap.Push(&q, &FANode{
						dist: dist[neig],
						v:    neig,
					})
				}
				res[index] = true
			}
		}
	}
	// fmt.Println(visited)
	return res
}
