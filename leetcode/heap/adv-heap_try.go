package heap

import "fmt"

type Point struct {
	x int
	y int
}

type Pointers []Point

func NewPointer(i, j int) Point {
	return Point{i, j}
}

func (h Pointers) Len() int          { return len(h) }
func (h Pointers) Comp(i, j int) int { return h[i].x - h[j].x }
func (h Pointers) Swap(i, j int)     { h[i], h[j] = h[j], h[i] }

func (h *Pointers) Push(a any) {
	*h = append(*h, a.(Point))
}

func (h *Pointers) Pop() any {
	x := (*h)[0]
	*h = (*h)[1:]
	return x
}

func Adv_Heap_Test() {
	p := Pointers{}

	Push(&p, NewPointer(1, 2))
	Push(&p, NewPointer(0, 2))
	Push(&p, NewPointer(3, 2))
	fmt.Println(p.Pop())
	fmt.Println(p.Pop())
	fmt.Println(p.Pop())
	fmt.Println(p.Pop())
	fmt.Println(p)
}
