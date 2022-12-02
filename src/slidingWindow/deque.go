package main

type Deque struct {
	nums []int
}

func NewDeque() *Deque {
	return new(Deque)
}

func (d *Deque) PushBack(n int) {
	d.nums = append(d.nums, n)
}

func (d *Deque) PushFront(n int) {
	d.nums = append([]int{n}, d.nums...)
}

func (d *Deque) PopBack() int {
	i := len(d.nums) - 1
	defer func() {
		d.nums = append(d.nums[:i], d.nums[i+1:]...)
	}()
	return d.nums[i]
}

func (d *Deque) PopFront() int {
	defer func() {
		d.nums = d.nums[1:]
	}()
	return d.nums[0]
}

func (d *Deque) Front() int {
	return d.nums[0]
}

func (d *Deque) Back() int {
	return d.nums[len(d.nums)-1]
}

func (d *Deque) Len() int {
	return len(d.nums)
}

func (d *Deque) IsEmpty() bool {
	return len(d.nums) == 0
}
