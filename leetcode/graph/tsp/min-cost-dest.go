package graph

//https://leetcode.com/problems/minimum-cost-to-reach-destination-in-time/solutions/3166087/21ms-beats-100-dijkstra-with-optimization/

// Diksjtra
func minCost(maxTime int, edges [][]int, passingFees []int) int {
	// setup graph
	g := map[int][][]int{}

	for _, v := range edges {
		if _, ok := g[v[0]]; !ok {
			g[v[0]] = [][]int{}
		}
		if _, ok := g[v[1]]; !ok {
			g[v[1]] = [][]int{}
		}
		g[v[0]] = append(g[v[0]], []int{v[1], v[2]})
		g[v[1]] = append(g[v[1]], []int{v[0], v[2]})
	}

	time := make([]int, len(passingFees))
	for i := range time {
		time[i] = maxTime + 1
	}

	pq := NewHeap(func(a, b []int) int {
		if a[1] == b[1] {
			return a[2] - b[2]
		} else {
			return a[1] - b[1]
		}
	})
	// [vertex, time, cost] // run dijstra on time and cost
	pq.Push([]int{0, 0, passingFees[0]})
	time[0] = 0
	dest := len(passingFees) - 1
	for pq.size > 0 {
		curr := pq.Pop()

		if curr[0] == dest {
			return curr[2]
		}
		for _, nei := range g[curr[0]] {
			if nei[1]+curr[1] < time[nei[0]] {
				time[nei[0]] = nei[1] + curr[1]
				pq.Push([]int{nei[0], nei[1] + curr[1], curr[2] + passingFees[nei[0]]})
			}
		}
	}

	return -1
}

// TSP
func minCost_DP(maxTime int, edges [][]int, passingFees []int) int {
	// setup graph
	g := map[int][][]int{}

	for _, v := range edges {
		if _, ok := g[v[0]]; !ok {
			g[v[0]] = [][]int{}
		}
		if _, ok := g[v[1]]; !ok {
			g[v[1]] = [][]int{}
		}
		g[v[0]] = append(g[v[0]], []int{v[1], v[2]})
		g[v[1]] = append(g[v[1]], []int{v[0], v[2]})
	}

	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var M int = 1e9
	dest := len(passingFees) - 1
	memoise := map[[2]int]int{}
	var dp func(int, int) int
	dp = func(time, v int) int {

		if time < 0 {
			return M
		}

		if v == dest {
			if time >= 0 {
				return passingFees[v]
			} else {
				return M
			}
		}

		if v, ok := memoise[[2]int{time, v}]; ok {
			return v
		}

		ans := M
		for _, n := range g[v] {
			ans = Min(ans, passingFees[v]+dp(time-n[1], n[0]))
		}
		memoise[[2]int{time, v}] = ans
		return ans
	}

	res := dp(maxTime, 0)
	if res == M {
		return -1
	} else {
		return res
	}
}

// / Heap for Priority Queue
type Heap struct {
	size int
	heap [][]int
	comp func(a, b []int) int
}

func NewHeap(comp func(i, j []int) int) *Heap {
	return &Heap{
		size: 0,
		heap: [][]int{},
		comp: comp,
	}
}

func (h *Heap) Push(data []int) {
	h.heap = append(h.heap, data)
	h.size++
	percolateUp(len(h.heap)-1, h)
}

func (h *Heap) Pop() []int {
	if h.Size() > 0 {
		data := h.heap[0]
		swap(0, len(h.heap)-1, &h.heap)
		h.heap = h.heap[:len(h.heap)-1]
		h.size--
		percolateDown(0, h)
		return data
	}
	panic("Heap is Empty!")
}

func (h *Heap) Size() int {
	return h.size
}

func (h *Heap) Peek() []int {
	if len(h.heap) > 0 {
		return h.heap[0]
	}
	panic("Heap is Empty!")
}

func percolateDown(i int, h *Heap) {
	leftChild := (2 * i) + 1
	rightChild := (2 * i) + 2

	parent := i

	if leftChild < len(h.heap) && h.comp(h.heap[i], h.heap[leftChild]) > 0 {
		i = leftChild
	}

	if rightChild < len(h.heap) && h.comp(h.heap[i], h.heap[rightChild]) > 0 {
		i = rightChild
	}

	if parent != i {
		swap(parent, i, &h.heap)
		percolateDown(i, h)
	}
}

func percolateUp(i int, h *Heap) {
	parent := (i - 1) / 2
	if parent >= 0 {
		if h.comp(h.heap[parent], h.heap[i]) > 0 {
			swap(parent, i, &h.heap)
			percolateUp(parent, h)
		}
	}
}

func swap(i, j int, arr *[][]int) {
	(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
}
