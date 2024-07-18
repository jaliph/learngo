package trie

// https://leetcode.com/problems/implement-trie-prefix-tree/

type Node struct {
	node [26]*Node
	end  bool
}

type Trie struct {
	root *Node
}

func Constructor2() Trie {
	return Trie{
		root: &Node{},
	}
}

func (this *Trie) Insert(word string) {
	n := this.root
	for _, ch := range word {
		d := ch - 'a'
		if n.node[d] == nil {
			n.node[d] = &Node{}
		}
		n = n.node[d]
	}

	n.end = true
}

func (this *Trie) Traverse(word string) *Node {
	n := this.root
	for _, ch := range word {
		d := ch - 'a'
		if n.node[d] == nil {
			return n.node[d]
		}
		n = n.node[d]
	}
	return n
}

func (this *Trie) Search(word string) bool {
	n := this.Traverse(word)
	return n != nil && n.end
}

func (this *Trie) StartsWith(prefix string) bool {
	n := this.Traverse(prefix)
	return n != nil
}
