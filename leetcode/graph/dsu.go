package graph

type DSU struct {
	size       int
	parents    []int
	ranks      []int
	components int
}

func NewDSU(n int) *DSU {
	d := &DSU{
		size:       n,
		parents:    make([]int, n),
		ranks:      make([]int, n),
		components: n,
	}

	for i := 0; i < n; i++ {
		d.parents[i] = -1
	}
	for i := 0; i < n; i++ {
		d.ranks[i] = 1
	}

	return d
}

func (d *DSU) Find(i int) int {
	if d.parents[i] == -1 {
		return i
	}

	d.parents[i] = d.Find(d.parents[i])
	return d.parents[i]
}

func (d *DSU) Union(i, j int) bool {
	p1 := d.Find(i)
	p2 := d.Find(j)

	if p1 == p2 {
		return false
	} else {
		r1 := d.ranks[p1]
		r2 := d.ranks[p2]
		if r1 > r2 {
			d.parents[p2] = p1
			d.ranks[r1] += d.ranks[r2]
		} else {
			d.parents[p1] = p2
			d.ranks[r2] += d.ranks[r1]
		}
		d.components--
		return true
	}
}

func (d *DSU) IsConnected() bool {
	return d.components == 1
}
