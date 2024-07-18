package dp

// https://leetcode.com/problems/extra-characters-in-a-string/description/

// setting up trie
type Node struct {
	node [26]*Node
	end  bool
}

type Trie struct {
	root *Node
}

func (t *Trie) AddWord(s string) {
	n := t.root
	for _, ch := range s {
		d := ch - 'a'
		if n.node[d] == nil {
			n.node[d] = &Node{}
		}
		n = n.node[d]
	}
	n.end = true
}

func MinExtraChar(s string, dictionary []string) int {

	dict := map[string]struct{}{}
	for _, str := range dictionary {
		dict[str] = struct{}{}
	}

	NewTrie := func(dict []string) *Trie {
		trie := &Trie{
			root: &Node{},
		}
		for _, s := range dict {
			trie.AddWord(s)
		}
		return trie
	}

	trie := NewTrie(dictionary)

	// setting up memo
	n := len(s)
	memoise := make([]int, len(s)+1)
	for i := range memoise {
		memoise[i] = n + 1
	}
	memoise[len(s)] = 0

	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var recur func(int) int
	recur = func(idx int) int {
		if idx == len(s) {
			return 0
		}
		if memoise[idx] != n+1 {
			return memoise[idx]
		}

		n := trie.root
		res := 1 + recur(idx+1)
		for j := idx; j < len(s); j++ {
			d := s[j]
			if n.node[d] == nil {
				break
			}
			n = n.node[d]
			if n.end {
				res = Min(res, recur(j+1))
			}
		}
		memoise[idx] = res
		return res
	}
	return recur(0)
}

func MinExtraChar_1(s string, dictionary []string) int {

	dict := map[string]struct{}{}
	for _, str := range dictionary {
		dict[str] = struct{}{}
	}
	// setting up memo
	n := len(s)
	memoise := make([]int, len(s)+1)
	for i := range memoise {
		memoise[i] = n + 1
	}
	memoise[len(s)] = 0

	Min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var recur func(int) int
	recur = func(idx int) int {
		if idx == len(s) {
			return 0
		}
		if memoise[idx] != n+1 {
			return memoise[idx]
		}

		res := 1 + recur(idx+1)
		for j := idx; j < len(s); j++ {
			if _, ok := dict[s[idx:j+1]]; ok {
				res = Min(res, recur(j+1))
			}
		}
		memoise[idx] = res
		return res
	}
	return recur(0)
}
