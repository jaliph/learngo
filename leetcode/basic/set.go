package basic

type Set struct {
	items map[int]struct{}
}

func NewSet() *Set {
	return &Set{
		items: make(map[int]struct{}),
	}
}

func (set *Set) Add(val int) {
	set.items[val] = struct{}{}
}

func (set Set) Has(val int) bool {
	_, ok := set.items[val]
	return ok
}

func (set *Set) Remove(val int) {
	delete(set.items, val)
}

func (set Set) Size() int {
	return len(set.items)
}
