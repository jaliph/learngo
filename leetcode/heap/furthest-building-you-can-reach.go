package heap

import "container/heap"

func FurthestBuilding(heights []int, bricks int, ladders int) int {
	h := &IntHeap{}
	heap.Init(h)
	for i := 1; i < len(heights); i++ {
		diff := heights[i] - heights[i-1]
		if diff > 0 {
			heap.Push(h, diff)
			if h.Len() > ladders {
				bricks -= heap.Pop(h).(int)
			}
			if bricks < 0 {
				return i - 1
			}
		}
	}
	return len(heights) - 1
}
