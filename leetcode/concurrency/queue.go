package concurrency

type Queue struct {
	items chan interface{}
}

func NewQueue(size int) *Queue {
	return &Queue{
		items: make(chan interface{}, size),
	}
}

func (q *Queue) Push(item interface{}) {
	select {
	case q.items <- item:
	default:
		panic("Queue Full")
	}
}

func (q *Queue) Pop() interface{} {
	select {
	case e := <-q.items:
		return e
	default:
		panic("Queue Empty")
	}
}

func (q *Queue) Size() int {
	return len(q.items)
}
