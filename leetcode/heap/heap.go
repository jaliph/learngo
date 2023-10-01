package heap

type Heap struct {
	size int
	heap []int
	comp func(a, b int) int
}

func NewHeap(comp func(i, j int) int) *Heap {
	if comp == nil {
		c := func(a, b int) int {
			return a - b
		}
		return &Heap{
			size: 0,
			heap: []int{},
			comp: c,
		}
	}
	return &Heap{
		size: 0,
		heap: []int{},
		comp: comp,
	}
}

func (h *Heap) Push(data int) {
	h.heap = append(h.heap, data)
	h.size++
	percolateUp(len(h.heap)-1, h)
}

func (h *Heap) Pop() int {
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

func (h *Heap) Peek() int {
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

func swap(i, j int, arr *[]int) {
	(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
}
