package query

type Fenwick struct {
	size int
	fn   []int
}

func NewFenwick(size int) *Fenwick {
	return &Fenwick{size + 1, make([]int, size+1)}
}

func (f *Fenwick) Add(i, val int) {
	i++ // 1 based index
	for i < f.size {
		f.fn[i] += val
		i += (i & -i)
	}
}

func (f *Fenwick) sum(idx int) int {
	idx++ // 1 based index
	sum := 0
	for idx > 0 {
		sum += f.fn[idx]
		idx -= (idx & -idx)
	}
	return sum
}

func (f *Fenwick) Sum(l, r int) int {
	return f.sum(r) - f.sum(l-1)
}
