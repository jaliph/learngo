package heap

type HeapInterface interface {
	Len() int
	Comp(i, j int) int
	Swap(i, j int)
	Push(x any)
	Pop() any
}

func Push(h HeapInterface, a any) {
	h.Push(a)
	PercolateUp(h.Len()-1, h)
}

func Pop(h HeapInterface) (a any) {
	a = h.Pop()
	PercolateDown(0, h)
	return
}

func PercolateDown(i int, h HeapInterface) {
	leftChild := (2 * i) + 1
	rightChild := (2 * i) + 2

	parent := i

	if leftChild < h.Len() && h.Comp(i, leftChild) > 0 {
		i = leftChild
	}

	if rightChild < h.Len() && h.Comp(i, rightChild) > 0 {
		i = rightChild
	}

	if parent != i {
		h.Swap(parent, i)
		PercolateDown(i, h)
	}
}

func PercolateUp(i int, h HeapInterface) {
	parent := (i - 1) / 2
	if parent >= 0 {
		if h.Comp(parent, i) > 0 {
			h.Swap(parent, i)
			PercolateUp(parent, h)
		}
	}
}
