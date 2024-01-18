package basic

// https://leetcode.com/problems/design-hashset/
import "math"

type Node struct {
	data int
	next *Node
}

func NewNode(val int) *Node {
	return &Node{
		val,
		nil,
	}
}

type MyHashSet struct {
	set []*Node
}

func Constructor3() MyHashSet {
	return MyHashSet{
		set: make([]*Node, int(math.Pow(10, 4))),
	}
}

func (this *MyHashSet) Add(key int) {
	h := hash(key)
	if this.set[h] == nil {
		this.set[h] = NewNode(-1)
	}
	temp := this.set[h]
	for temp.next != nil {
		if temp.next.data == key {
			return
		}
		temp = temp.next
	}
	temp.next = NewNode(key)
}

func (this *MyHashSet) Remove(key int) {
	h := hash(key)
	temp := this.set[h]

	for temp != nil && temp.next != nil {
		if temp.next.data == key {
			temp.next = temp.next.next
			return
		}
		temp = temp.next
	}
}

func (this *MyHashSet) Contains(key int) bool {
	h := hash(key)
	temp := this.set[h]

	for temp != nil && temp.next != nil {
		if temp.next.data == key {
			return true
		}
		temp = temp.next
	}
	return false
}

func hash(key int) int {
	return key % int(math.Pow(10, 4))
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
