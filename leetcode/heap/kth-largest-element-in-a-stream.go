package heap

// https://leetcode.com/problems/kth-largest-element-in-a-stream
type KthLargest struct {
	sorted []int
	k      int
}

func NewKthLargest(k int, nums []int) KthLargest {
	kth := KthLargest{
		sorted: []int{},
		k:      k,
	}

	for _, n := range nums {
		kth.Insert(n)
		if len(kth.sorted) > k {
			kth.sorted = kth.sorted[:k]
		}
	}
	return kth
}

func (this *KthLargest) Insert(i int) {
	l := 0
	r := len(this.sorted)
	for l < r {
		mid := l + ((r - l) / 2)
		if this.sorted[mid] > i {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if l == len(this.sorted) {
		this.sorted = append(this.sorted, i)
	} else {
		this.sorted = append(this.sorted[:l+1], this.sorted[l:]...)
		this.sorted[l] = i
	}
}

func (this *KthLargest) Add(val int) int {
	this.Insert(val)
	if len(this.sorted) > this.k {
		this.sorted = this.sorted[:this.k]
	}
	return this.sorted[this.k-1]
}
