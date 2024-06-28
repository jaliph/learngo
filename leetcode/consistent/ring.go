package consistent

import (
	"hash/crc32"
	"sort"
)

// Ring is a network of distributed nodes.
type Ring struct {
	Nodes Nodes
}                                  // Nodes is an array of nodes.
type Nodes []*Node                 // Node is a single entity in a ring.
func (n Nodes) Len() int           { return len(n) }
func (n Nodes) Less(i, j int) bool { return n[i].HashId < n[j].HashId }
func (n Nodes) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }

type Node struct {
	Id     string
	HashId uint32
}

// Initializes new distribute network of nodes or a ring.

func (r *Ring) AddNode(id string) {
	node := NewNode(id)
	r.Nodes = append(r.Nodes, node)
	sort.Sort(r.Nodes)
}

// Removes node from the ring if it exists, else returns
// ErrNodeNotFound.
func (r *Ring) RemoveNode(id string) error

// Gets node which is mapped to the key. Return value is identifer
// of the node given in `AddNode`.
func (r *Ring) Get(key string) string {
	searchfn := func(i int) bool {
		return r.Nodes[i].HashId >= crc32.Checksum([]byte(key), crc32.IEEETable)
	}
	i := sort.Search(r.Nodes.Len(), searchfn)
	if i >= r.Nodes.Len() {
		i = 0
	}
	return r.Nodes[i].Id
}

// Adds node to the ring.
func NewRing() *Ring {
	return &Ring{Nodes: Nodes{}}
}

func NewNode(id string) *Node {
	return &Node{
		Id:     id,
		HashId: crc32.Checksum([]byte(id), crc32.IEEETable),
	}
}
