package heap

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type ListNodeHeap []*ListNode

func (l ListNodeHeap) Less(i, j int) bool { return l[i].Val < l[j].Val }
func (l ListNodeHeap) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ListNodeHeap) Len() int           { return len(l) }
func (l *ListNodeHeap) Pop() any {
	old := *l
	node := old[len(old)-1]
	*l = old[:len(old)-1]
	return node
}
func (l *ListNodeHeap) Push(x any) {
	*l = append(*l, x.(*ListNode))
}

func MergeKLists(lists []*ListNode) *ListNode {
	h := &ListNodeHeap{}
	heap.Init(h)

	for _, l := range lists {
		if l != nil {
			heap.Push(h, l)
		}
	}

	head := &ListNode{}
	w := head
	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
		w.Next = node
		w = w.Next
	}

	return head.Next
}
