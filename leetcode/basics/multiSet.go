package basic

type MultiSet struct {
	items map[int]int
}

func NewMultiSet() *MultiSet {
	return &MultiSet{
		items: make(map[int]int),
	}
}

func (s *MultiSet) Add(val int) {
	s.items[val]++
}

func (s *MultiSet) Delete(val int) {
	delete(s.items, val)
}

func (s *MultiSet) Has(val int) bool {
	_, ok := s.items[val]
	return ok
}

func (s *MultiSet) Size() int {
	return len(s.items)
}
