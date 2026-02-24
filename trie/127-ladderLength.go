package week07

// Two directions BFS
// It is well than single direction BFS
func ladderLength(beginWord string, endWord string, wordList []string) int {
	if wordList == nil || len(wordList) == 0 {
		return 0
	}
	if endWord == "" || beginWord == "" {
		return 0
	}

	myMap := make(map[string]bool, len(wordList))
	for _, str := range wordList {
		myMap[str] = true
	}
	if _, ok := myMap[endWord]; !ok {
		return 0
	}
	delete(myMap, beginWord)

	visited := make(map[string]bool, len(wordList))
	visited[beginWord] = true
	beginVisited, endVisited := map[string]bool{beginWord: true}, map[string]bool{endWord: true}
	step := 1

	for len(beginVisited) > 0 && len(endVisited) > 0 {
		if len(beginVisited) > len(endVisited) {
			beginVisited, endVisited = endVisited, beginVisited
		}

		newVisited := map[string]bool{}
		for word := range beginVisited {
			for i := 0; i < len(word); i++ {
				for ch := 'a'; ch <= 'z'; ch++ {
					if rune(word[i]) == ch {
						continue
					}
					newWord := word[:i] + string(ch) + word[i+1:]
					if _, ok := myMap[newWord]; ok {
						if _, ok := endVisited[newWord]; ok {
							step++
							return step
						}

						if _, yes := visited[newWord]; !yes {
							visited[newWord] = true
							newVisited[newWord] = true
						}
					}
				}
			}
		}
		beginVisited = newVisited
		step++
	}

	return 0
}
