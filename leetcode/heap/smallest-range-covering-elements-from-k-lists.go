package heap

import (
	"container/heap"
)

// https://leetcode.com/problems/smallest-range-covering-elements-from-k-lists/

type Entry struct {
	val int
	idx int
	pos int
}

type Ranges []Entry

func (r Ranges) Len() int           { return len(r) }
func (r Ranges) Less(i, j int) bool { return r[i].val < r[j].val }
func (r Ranges) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r *Ranges) Push(x any) {
	*r = append(*r, x.(Entry))
}
func (r *Ranges) Pop() any {
	old := *r
	l := len(old)
	data := old[l-1]
	*r = old[:l-1]
	return data
}

func SmallestRange(nums [][]int) []int {
	const INF int = 1e9
	rangeMax := INF
	rangeMin := -INF
	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	max := -INF

	r := &Ranges{}
	heap.Init(r)
	for i, vals := range nums {
		heap.Push(r, Entry{vals[0], i, 0})
		max = Max(max, vals[0])
	}

	for len(*r) == len(nums) {
		pop := heap.Pop(r).(Entry)
		rangeCurr := max - pop.val + 1
		if rangeCurr < rangeMax-rangeMin+1 {
			rangeMax = max
			rangeMin = pop.val
		}

		if pop.pos+1 < len(nums[pop.idx]) {
			heap.Push(r, Entry{nums[pop.idx][pop.pos+1], pop.idx, pop.pos + 1})
			max = Max(max, nums[pop.idx][pop.pos+1])
		}
	}

	// fmt.Println(*r)

	return []int{rangeMin, rangeMax}
}
