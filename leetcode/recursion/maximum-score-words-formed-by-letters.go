package recursion

// https://leetcode.com/problems/maximum-score-words-formed-by-letters/

func MaxScoreWords(words []string, letters []byte, score []int) int {

	countMap := map[byte]int{}

	copyMap := func(m map[byte]int) map[byte]int {
		n := map[byte]int{}
		for k, v := range m {
			n[k] = v
		}
		return n
	}

	Max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for _, letter := range letters {
		countMap[letter-'a']++
	}

	var dfs func(int, map[byte]int) int
	dfs = func(idx int, m map[byte]int) int {
		if idx == len(words) {
			return 0
		}

		// dont take
		dontTake := dfs(idx+1, m)

		// take
		word := words[idx]
		take := 0
		canTake := true
		newMap := copyMap(m)
		currentScore := 0
		for i := range word {
			d := word[i] - 'a'
			if newMap[d] <= 0 {
				canTake = false
				break
			}
			currentScore += score[d]
			newMap[d]--
		}
		if canTake {
			take = currentScore + dfs(idx+1, newMap)
		}

		return Max(take, dontTake)

	}
	return dfs(0, countMap)
}

// func main() {

// 	fmt.Println(recursion.MaxScoreWords(
// 		[]string{"dog", "cat", "dad", "good"}, []byte{'a', 'a', 'c', 'd', 'd', 'd', 'g', 'o', 'o'}, []int{1, 0, 9, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
// }
