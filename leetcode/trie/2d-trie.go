package trie

// https://leetcode.com/problems/implement-trie-prefix-tree

import "fmt"

const TRIE_SIZE = 10

type Trie2D struct {
	ptr   int
	nodes [][]int
}

func NewTrieNode() *Trie2D {
	c := 27 // columns
	s := &Trie2D{
		ptr:   0,
		nodes: make([][]int, TRIE_SIZE),
	}
	rows := make([]int, TRIE_SIZE*c)

	for i := range s.nodes {
		s.nodes[i] = rows[i*c : (i+1)*c]
	}
	return s
}

func (s *Trie2D) AddWord(str string) {
	curr := 0
	for _, v := range str {
		d := int(v - 'a')
		if s.nodes[curr][d] == 0 {
			s.ptr++
			s.nodes[curr][d] = s.ptr
		}
		curr = s.nodes[curr][d]
	}
	s.nodes[curr][26] = 1
}

func (s *Trie2D) Search(str string) bool {
	curr := 0
	for _, v := range str {
		d := int(v - 'a')
		if s.nodes[curr][d] == 0 {
			return false
		}
		curr = s.nodes[curr][d]
	}
	return s.nodes[curr][26] == 1
}

func (s *Trie2D) StartsWith(str string) bool {
	curr := 0
	for _, v := range str {
		d := int(v - 'a')
		if s.nodes[curr][d] == 0 {
			return false
		}
		curr = s.nodes[curr][d]
	}
	for i := 0; i < 26; i++ {
		if s.nodes[curr][i] != 0 {
			return true
		}
	}
	return s.nodes[curr][26] == 1
}

func Driver_Trie2D() {
	s := NewTrieNode()
	s.AddWord("abc")
	s.AddWord("adf")
	s.AddWord("bad")
	Print2D(s.nodes)
	str := "adf"
	fmt.Println("str is present ?", s.Search(str))
	str = "bad"
	fmt.Println("str startswith ?", s.StartsWith(str))
}

func Print2D(grid [][]int) {
	for r, _ := range grid {
		for _, colValue := range grid[r] {
			fmt.Print(colValue, " ")
		}
		fmt.Println()
	}
}
