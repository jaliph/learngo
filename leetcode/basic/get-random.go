package basic

import (
	"math/rand"
)

// https://leetcode.com/problems/insert-delete-getrandom-o1

type RandomizedSet struct {
	vMap     map[int]int
	list2map []int
}

func ConstructorRS() RandomizedSet {
	return RandomizedSet{
		vMap:     make(map[int]int),
		list2map: []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.vMap[val]; ok {
		return !ok
	}
	this.list2map = append(this.list2map, val)
	this.vMap[val] = len(this.list2map) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if toDelIdx, ok := this.vMap[val]; ok {
		// move the idx value to the last and delete from the map
		this.list2map[toDelIdx] = this.list2map[len(this.list2map)-1]
		this.vMap[this.list2map[toDelIdx]] = toDelIdx

		this.list2map = this.list2map[:len(this.list2map)-1]
		delete(this.vMap, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	idx := rand.Intn(len(this.list2map))
	return this.list2map[idx]
}
