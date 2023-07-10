package main

import "fmt"

func main() {
	// h := heap{
	// 	store: []int{0, 1, 5, 6, 3, 2, 4, 5},
	// }

	h2 := heap{
		store: []int{},
	}

	// // // h2.store[1] = 4
	fmt.Println(h2)
	h2.insertValue(10)
	h2.insertValue(11)
	// h2.insertValue(2)
	// h2.insertValue(22)
	// h2.insertValue(23)
	// fmt.Println(h2)
	// fmt.Println(h2.maximum())

	// h2.store = h2.store[1:3]
	// fmt.Println(h2)
	// fmt.Println(len(h2.store))
	// fmt.Println(cap(h2.store))
	fmt.Println(h2)
	// h.build_maxheap()
	// fmt.Println(h.store)
	// h.build_minheap()
	// fmt.Println(h.store)
}

type heap struct {
	store []int
}

// gives the maxmimum element from the heap
func (h *heap) maximum() int {
	if len(h.store) < 2 {
		return -1
	}
	return h.store[1]
}

// returns and removes the maximum
func (h *heap) extractMaximum() int {
	if len(h.store) < 2 {
		fmt.Println("heap is empty")
		return -1
	}
	max := h.store[1]
	h.store[1] = h.store[len(h.store)-1]
	h.store = h.store[:len(h.store)-1]
	h.max_heapify(1)
	return max
}

// maxHeap
func (h *heap) setValueAtPos(pos int, data int) {

	if data < h.store[pos] {
		//not possible
		fmt.Println("not a possible heap solution")
		return
	}
	h.store[pos] = data
	// pos/2 is parent of pos in the array
	// then check for the next parent at pos/2 till we reach to absolute parent at 1
	for pos > 1 && h.store[pos/2] < h.store[pos] {
		h.store[pos/2], h.store[pos] = h.store[pos], h.store[pos/2]
		pos /= 2
	}
}

func (h *heap) insertValue(data int) {
	h.store = append(h.store, -1) // increase the length of the heap
	if len(h.store) == 1 {
		h.store = append(h.store, -1)
	}
	h.setValueAtPos(len(h.store)-1, data)
}

func (h *heap) build_maxheap() {
	for i := (len(h.store) - 1) / 2; i >= 1; i-- {
		h.max_heapify(i)
	}
}

func (h *heap) build_minheap() {
	for i := (len(h.store) - 1) / 2; i >= 1; i-- {
		h.min_heapify(i)
	}
}

// Process and Convert the array to HEAP
func (h *heap) max_heapify(i int) {
	var largest int
	left := i * 2
	right := i*2 + 1
	fmt.Println("left is ", left)
	fmt.Println("right is ", right)

	if left < len(h.store) && h.store[left] > h.store[i] {
		largest = left
		fmt.Println("left is largest")
	} else {
		fmt.Println("none is largest")
		largest = i
	}
	if right < len(h.store) && h.store[right] > h.store[largest] {
		largest = right
		fmt.Println("right is largest")
	}
	if largest != i {
		h.store[i], h.store[largest] = h.store[largest], h.store[i]
		h.max_heapify(largest)
	}
}

func (h *heap) min_heapify(i int) {
	var smallest int
	left := i * 2
	right := i*2 + 1
	if left < len(h.store) && h.store[left] < h.store[i] {
		smallest = left
	} else {
		smallest = i
	}
	if right < len(h.store) && h.store[right] < h.store[smallest] {
		smallest = right
	}
	if smallest != i {
		h.store[i], h.store[smallest] = h.store[smallest], h.store[i]
		h.min_heapify(smallest)
	}
}
