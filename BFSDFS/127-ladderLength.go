package week04

import "container/list"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if wordList == nil {
		return 0
	}

	// use map to store all words
	myMap := make(map[string]bool, len(wordList))
	for _, word := range wordList {
		myMap[word] = true
	}
	if _, ok := myMap[endWord]; !ok {
		return 0
	}
	delete(myMap, beginWord)

	// init and tag words had been visited
	visited := make(map[string]bool, len(wordList))
	visited[beginWord] = true
	myQue := list.New()
	myQue.PushBack(beginWord)
	step := 1

	// BFS
	for myQue.Len() > 0 {
		// should not use myQue.Len() directly, it's dymatic
		l := myQue.Len()
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
						// ok
						if newWord == endWord {
							return step + 1
						}
						// mark the word
						if _, ok := visited[newWord]; !ok {
							visited[newWord] = true
							myQue.PushBack(newWord)
						}
					}
				}
			}
		}
		step++
	}

	return 0
}
