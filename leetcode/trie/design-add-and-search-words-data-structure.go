package trie

import "fmt"

//	type Node struct {
//		node [26]*Node
//		end  bool
//	}

type WordDictionary struct {
	root *Node
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: &Node{},
	}
}

func (this *WordDictionary) AddWord(word string) {
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

func (this *WordDictionary) SearchIndex(idx int, n *Node, word string) bool {
	for i := idx; i < len(word); i++ {
		if word[i] == '.' {
			for _, node := range n.node {
				if node != nil && this.SearchIndex(i+1, node, word) {
					return true
				}
			}
			return false
		} else {
			d := word[i] - 'a'
			if n.node[d] == nil {
				return false
			}
			n = n.node[d]
		}
	}
	return n.end
}

func (this *WordDictionary) Search(word string) bool {
	return this.SearchIndex(0, this.root, word)
}

func Driver() {
	obj := Constructor()
	obj.AddWord("bad")
	obj.AddWord("add")
	val1 := obj.Search("a.e")
	fmt.Println(val1)
}
