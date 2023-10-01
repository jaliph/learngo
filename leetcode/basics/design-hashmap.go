package basic

import "math"

type HashNode struct {
	key  int
	data int
	next *HashNode
}

func NewHashNode(key, val int) *HashNode {
	return &HashNode{
		key:  key,
		data: val,
		next: nil,
	}
}

type MyHashMap struct {
	hash []*HashNode
}

func Constructor() MyHashMap {
	return MyHashMap{
		hash: make([]*HashNode, int(math.Pow(10, 1))),
	}
}

func (this *MyHashMap) Put(key int, value int) {
	h := hasher(key)

	if this.hash[h] == nil {
		this.hash[h] = NewHashNode(-1, -1)
	}

	temp := this.hash[h]
	for temp.next != nil {
		if temp.next.key == key {
			temp.next.data = value
			return
		}
		temp = temp.next
	}

	temp.next = NewHashNode(key, value)
}

func (this *MyHashMap) Get(key int) int {
	h := hasher(key)
	temp := this.hash[h]

	for temp != nil && temp.next != nil {
		if temp.next.key == key {
			return temp.next.data
		}
		temp = temp.next
	}

	return -1
}

func (this *MyHashMap) Remove(key int) {
	h := hasher(key)
	temp := this.hash[h]

	for temp != nil && temp.next != nil {
		if temp.next.key == key {
			temp.next = temp.next.next
		}
		temp = temp.next
	}
}

func hasher(key int) int {
	return key % int(math.Pow(10, 1))
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
