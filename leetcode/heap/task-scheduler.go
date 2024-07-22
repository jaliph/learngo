package heap

import "container/heap"

// https://leetcode.com/problems/task-scheduler/
func LeastInterval(tasks []byte, n int) int {

	freq := map[byte]int{}
	for _, v := range tasks {
		freq[v]++
	}

	h := &IntHeap{}
	heap.Init(h)
	for _, v := range freq {
		heap.Push(h, v)
	}

	q := [][2]int{}
	time := 0

	for h.Len() > 0 || len(q) > 0 {
		time++
		curr := heap.Pop(h).(int)
		curr--

		if curr > 0 {
			q = append(q, [2]int{curr, time + n})
		}

		for len(q) > 0 && q[0][1] <= time {
			pop := q[0]
			q = q[1:]
			heap.Push(h, pop[0])
		}
	}

	return time
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}
