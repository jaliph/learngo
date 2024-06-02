package heap

import (
	"container/heap"
	"fmt"
)

type Item struct {
	data     int
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// pop the highest first
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	item := old[len(old)-1]
	item.index = -1 // for safety
	*pq = old[:len(old)-1]
	return item
}

func (pq *PriorityQueue) Update(item *Item, value int, priority int) {
	item.data = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func Example_Driver() {
	items := map[int]int{
		1: 100,
		2: 300,
		3: 400,
	}

	pq := PriorityQueue{}
	i := 0
	for k, v := range items {
		pq = append(pq, &Item{
			data:     k,
			priority: v,
			index:    i,
		})
		i++
	}

	heap.Init(&pq)

	fmt.Println(pq)
	fmt.Println(heap.Pop(&pq))
	heap.Push(&pq, &Item{
		data:     4,
		priority: 1000,
	})
	fmt.Println(heap.Pop(&pq))
}
