package heap

import "fmt"

// https://leetcode.com/problems/seat-reservation-manager/

type SeatManager struct {
	seats []int
}

func Constructor(n int) SeatManager {
	sm := SeatManager{make([]int, n)}

	for i := 1; i <= n; i++ {
		sm.seats[i-1] = i
	}
	return sm
}

func (this *SeatManager) Reserve() int {
	seat := this.seats[0]
	this.seats = this.seats[1:]
	return seat
}

func (this *SeatManager) Unreserve(seatNumber int) {
	BinaryInsertion(&this.seats, seatNumber)
}

func BinaryInsertion(arr *[]int, target int) {
	l := 0
	r := len(*arr)
	var mid int
	for l < r {
		mid = l + (r-l)/2
		if target > (*arr)[mid] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	fmt.Println(l)
	if l == len(*arr) {
		*arr = append(*arr, target)
	} else {
		*arr = append((*arr)[:l+1], (*arr)[l:]...)
		(*arr)[l] = target
	}
}

/** Heap Way
type SeatManager struct {
	seats *Heap
}

func Constructor(n int) SeatManager {
	sm := SeatManager{
		seats: NewHeap(func(i, j int) int {
			return i - j
		}),
	}
	sm.seats.size = n
	for i := 1; i <= n; i++ {
		sm.seats.heap = append(sm.seats.heap, i)
	}
	return sm
}

func (this *SeatManager) Reserve() int {
	return this.seats.Pop()
}

func (this *SeatManager) Unreserve(seatNumber int) {
	this.seats.Push(seatNumber)
}
**/

// func Driver() {
// 	s := Constructor(5)

// 	fmt.Println(s.seats)
// 	fmt.Println(s.Reserve())
// 	fmt.Println(s.seats)
// 	s.Unreserve(5)
// 	fmt.Println(s.seats)
// }
