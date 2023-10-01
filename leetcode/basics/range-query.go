package basic

type NumArray struct {
	arr []int
}

func Constructor1(nums []int) NumArray {
	numArray := NumArray{
		arr: make([]int, len(nums)),
	}
	prefix := 0
	for i, v := range nums {
		numArray.arr[i] = prefix + v
		prefix = numArray.arr[i]
	}
	return numArray
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.arr[right]
	}
	return this.arr[right] - this.arr[left-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
