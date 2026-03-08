package week04

import "container/list"

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	if wordList == nil {
		return [][]string{}
	}

	// use map to store all words
	myMap := make(map[string]bool, len(wordList))
	for _, word := range wordList {
		myMap[word] = true
	}
	if _, ok := myMap[endWord]; !ok {
		return [][]string{}
	}
	delete(myMap, beginWord)

	// init and tag words had been visited
	visited := make(map[string]bool, len(wordList))
	visited[beginWord] = true
	myQue := list.New()
	myQue.PushBack(beginWord)
	found := false
	// map[string1][]{string11, string12}, string11 and stirng12 is the next word of string1
	relation := make(map[string][]string)

	// BFS
	for myQue.Len() > 0 {
		// should not use myQue.Len() directly, it's dymatic
		l := myQue.Len()
		nextVisted := map[string]bool{}

		for n := 0; n < l; n++ {
			word := myQue.Front().Value.(string)
			myQue.Remove(myQue.Front())
			// replace the char
			for i := 0; i < len(word); i++ {
				// only low char
				for ch := 'a'; ch <= 'z'; ch++ {
					if ch == rune(word[i]) {
						continue
					}
					newWord := word[:i] + string(ch) + word[i+1:]
					if exit, ok := myMap[newWord]; ok && exit {
						if _, ok := visited[newWord]; !ok {
							// found
							if newWord == endWord {
								found = true
							}
							// in order to avoid to push one word repeatly
							if _, ook := nextVisted[newWord]; !ook {
								myQue.PushBack(newWord)
								nextVisted[newWord] = true
							}

							// to make the relation of words
							if words, ook := relation[word]; ook {
								words := append(words, newWord)
								relation[word] = words
							} else {
								relation[word] = []string{newWord}
							}
						}
					}
				}
			}
		}
		if found {
			break
		}
		for w := range nextVisted {
			visited[w] = true
		}
	}

	return getWordsPathByDfs(beginWord, endWord, relation, []string{beginWord}, [][]string{})
}

// get path by dfs and recursion
func getWordsPathByDfs(beginWord, endWord string, relation map[string][]string, path []string, res [][]string) [][]string {
	if beginWord == endWord {
		res = append(res, path)
		return res
	}

	var words []string
	if _, ok := relation[beginWord]; ok {
		words = relation[beginWord]
	} else {
		return res
	}
	for _, word := range words {
		path = append(path, word)
		res = getWordsPathByDfs(word, endWord, relation, path, res)
		path = append([]string{}, path[:len(path)-1]...)
	}

	return res
}
