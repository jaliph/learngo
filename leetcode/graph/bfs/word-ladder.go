// https://leetcode.com/problems/word-ladder/

package graph

type Node struct {
	word string
	cost int
}

func NewNode(s string, c int) Node {
	return Node{s, c}
}

func LadderLength(beginWord string, endWord string, wordList []string) int {

	// check endword exists int he wordlist
	idx := len(wordList)
	for v, w := range wordList {
		if endWord == w {
			idx = v
			break
		}
	}

	if idx == len(wordList) {
		return 0
	}

	g := make(map[string][]string)

	linkCreator(beginWord, g)
	for _, w := range wordList {
		linkCreator(w, g)
	}

	l := func() int {

		q := []Node{}
		q = append(q, NewNode(endWord, 1))
		visited := make(map[string]bool)
		visited[endWord] = true

		for len(q) > 0 {
			curr := q[0]
			q = q[1:]

			if curr.word == beginWord {
				return curr.cost
			}

			for i := 0; i < len(curr.word); i++ {
				link := string(curr.word[0:i]) + "*" + string(curr.word[i+1:])

				for _, v := range g[link] {
					if _, ok := visited[v]; !ok {
						visited[v] = true
						q = append(q, NewNode(v, curr.cost+1))
					}
				}
			}
		}

		return 0
	}()

	return l
}

func linkCreator(word string, g map[string][]string) {
	for i := 0; i < len(word); i++ {
		link := string(word[0:i]) + "*" + string(word[i+1:])
		if _, ok := g[link]; !ok {
			g[link] = []string{}
		}
		g[link] = append(g[link], word)
	}
}
