package week04

import "container/list"

/*
127. Word Ladder
字典 wordList 中从单词 beginWord 到 endWord 的 转换序列 是一个按下述规格形成的序列 beginWord -> s1 -> s2 -> ... -> sk：

每一对相邻的单词只差一个字母。

	对于 1 <= i <= k 时，每个 si 都在 wordList 中。注意， beginWord 不需要在 wordList 中。

sk == endWord
给你两个单词 beginWord 和 endWord 和一个字典 wordList ，返回 从 beginWord 到 endWord 的 最短转换序列 中的 单词数目 。
如果不存在这样的转换序列，返回 0 。

示例 1：

输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
输出：5
解释：一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog", 返回它的长度 5。
示例 2：

输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
输出：0
解释：endWord "cog" 不在字典中，所以无法进行转换。
*/
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

func ladderLength1(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool, len(wordList))
	for _, word := range wordList {
		wordSet[word] = true
	}
	if _, ok := wordSet[endWord]; !ok {
		return 0
	}

	allWords := append([]string{beginWord}, wordList...)
	patternMap := make(map[string][]string)
	for _, word := range allWords {
		chars := []rune(word)
		for i := 0; i < len(chars); i++ {
			pattern := word[:i] + "*" + word[i+1:]
			patternMap[pattern] = append(patternMap[pattern], word)
		}
	}

	visited := make(map[string]struct{}, len(wordList))
	queue := []string{beginWord}
	pathLen := 1

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			word := queue[0]
			queue = queue[1:]

			chars := []rune(word)
			for j := 0; j < len(chars); j++ {
				pattern := word[:j] + "*" + word[j+1:]

				for _, neighbor := range patternMap[pattern] {
					if neighbor == word {
						continue
					}
					if neighbor == endWord {
						return pathLen + 1
					}
					if wordSet[neighbor] {
						if _, ok := visited[neighbor]; !ok {
							visited[neighbor] = struct{}{}
							queue = append(queue, neighbor)
						}
					}
				}
			}
		}
		pathLen++
	}

	return 0
}
