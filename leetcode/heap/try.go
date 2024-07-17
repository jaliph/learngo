package heap

import (
	"container/heap"
	"fmt"
)

type Ints []int

func (i Ints) Less(x, y int) bool {
	return i[x] < i[y]
}
func (i Ints) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}
func (i Ints) Len() int {
	return len(i)
}
func (i *Ints) Push(x any) {
	*i = append(*i, x.(int))
}
func (i *Ints) Pop() any {
	data := (*i)[len(*i)-1]
	*i = (*i)[:len(*i)-1]
	return data
}

func Try() {

	x := Ints{1, 2, 3}
	heap.Init(&x)
	heap.Push(&x, -4)
	fmt.Println(heap.Pop(&x))
	fmt.Println(heap.Pop(&x))
	fmt.Println(heap.Pop(&x))

}
