package week04

/*
126. Word Ladder II
按字典 wordList 完成从单词 beginWord 到单词 endWord 转化，一个表示此过程的 转换序列 是形式上像 beginWord -> s1 -> s2 -> ... -> sk 这样的单词序列，并满足：

每对相邻的单词之间仅有单个字母不同。
转换过程中的每个单词 si（1 <= i <= k）必须是字典 wordList 中的单词。注意，beginWord 不必是字典 wordList 中的单词。
sk == endWord
给你两个单词 beginWord 和 endWord ，以及一个字典 wordList 。请你找出并返回所有从 beginWord 到 endWord 的 最短转换序列 ，如果不存在这样的转换序列，返回一个空列表。每个序列都应该以单词列表 [beginWord, s1, s2, ..., sk] 的形式返回。

示例 1：

输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
输出：[["hit","hot","dot","dog","cog"],["hit","hot","lot","log","cog"]]
解释：存在 2 种最短的转换序列：
"hit" -> "hot" -> "dot" -> "dog" -> "cog"
"hit" -> "hot" -> "lot" -> "log" -> "cog"
示例 2：

输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
输出：[]
解释：endWord "cog" 不在字典 wordList 中，所以不存在符合要求的转换序列。
*/

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	// 1. 预处理：将单词列表转换为集合
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}

	// 检查endWord是否在单词列表中
	if !wordSet[endWord] {
		return [][]string{}
	}

	// 移除beginWord，避免重复访问
	delete(wordSet, beginWord)

	// 2. 使用双向BFS查找最短路径
	visited := make(map[string]int)      // 记录单词所在的层级
	parents := make(map[string][]string) // 记录父节点关系
	queue := []string{beginWord}
	visited[beginWord] = 1
	found := false
	level := 1

	for len(queue) > 0 && !found {
		// 记录当前层需要处理的单词数量
		levelSize := len(queue)

		// 处理当前层的所有单词
		for i := 0; i < levelSize; i++ {
			word := queue[0]
			queue = queue[1:]

			// 生成当前单词的所有可能的邻居
			chars := []rune(word)
			for j := 0; j < len(chars); j++ {
				originalChar := chars[j]

				// 尝试用'a'到'z'替换当前字符
				for c := 'a'; c <= 'z'; c++ {
					if c == originalChar {
						continue
					}

					chars[j] = c
					newWord := string(chars)

					// 如果新单词是目标单词
					if newWord == endWord {
						found = true
						// 记录父节点关系
						parents[newWord] = append(parents[newWord], word)
					} else if wordSet[newWord] {
						// 如果新单词在单词列表中
						if visitedLevel, ok := visited[newWord]; !ok {
							// 新单词第一次访问
							visited[newWord] = level + 1
							queue = append(queue, newWord)
							parents[newWord] = []string{word}
						} else if visitedLevel == level+1 {
							// 新单词在同一层级被访问，记录多条路径
							parents[newWord] = append(parents[newWord], word)
						}
					}
				}

				// 恢复原始字符
				chars[j] = originalChar
			}
		}

		level++

		// 如果找到了目标单词，处理完当前层就结束
		if found {
			// 不再添加新的节点到队列
			queue = nil
		}
	}

	// 3. 如果没有找到路径，返回空
	if !found {
		return [][]string{}
	}

	// 4. 使用DFS回溯构建所有路径
	result := [][]string{}

	// 从终点向起点回溯
	var dfs func(string, []string)
	dfs = func(currentWord string, currentPath []string) {
		// 将当前单词添加到路径开头
		currentPath = append([]string{currentWord}, currentPath...)

		// 如果到达起点，将路径添加到结果
		if currentWord == beginWord {
			result = append(result, currentPath)
			return
		}

		// 遍历所有父节点
		for _, parent := range parents[currentWord] {
			dfs(parent, currentPath)
		}
	}

	// 从终点开始回溯
	dfs(endWord, []string{})

	return result
}

func findLadders1(beginWord string, endWord string, wordList []string) [][]string {
	// 1. 检查endWord是否在单词列表中
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}

	if !wordSet[endWord] {
		return [][]string{}
	}

	// 2. 构建模式映射，加速邻居查找
	// 模式示例：对于"hot"，模式有："*ot", "h*t", "ho*"
	patternMap := make(map[string][]string)

	// 添加beginWord到单词列表，以便构建其模式
	allWords := append([]string{beginWord}, wordList...)

	for _, word := range allWords {
		chars := []rune(word)
		for i := 0; i < len(chars); i++ {
			// 创建模式
			pattern := word[:i] + "*" + word[i+1:]
			patternMap[pattern] = append(patternMap[pattern], word)
		}
	}

	// 3. 使用BFS查找最短路径
	visited := make(map[string]int)
	parents := make(map[string][]string)
	queue := []string{beginWord}
	visited[beginWord] = 1
	found := false

	for len(queue) > 0 && !found {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			word := queue[0]
			queue = queue[1:]

			// 生成所有可能的模式
			chars := []rune(word)
			for j := 0; j < len(chars); j++ {
				pattern := word[:j] + "*" + word[j+1:]

				// 查找所有与该模式匹配的单词
				for _, neighbor := range patternMap[pattern] {
					if neighbor == word {
						continue
					}

					if neighbor == endWord {
						found = true
						parents[neighbor] = append(parents[neighbor], word)
					} else if wordSet[neighbor] {
						if level, ok := visited[neighbor]; !ok {
							// 第一次访问
							visited[neighbor] = visited[word] + 1
							queue = append(queue, neighbor)
							parents[neighbor] = []string{word}
						} else if level == visited[word]+1 {
							// 同一层级再次访问
							parents[neighbor] = append(parents[neighbor], word)
						}
					}
				}
			}
		}

		if found {
			break
		}
	}

	// 4. 如果没有找到路径
	if !found {
		return [][]string{}
	}

	// 5. DFS回溯构建路径
	result := [][]string{}

	var backtrack func(string, []string)
	backtrack = func(currentWord string, path []string) {
		if currentWord == beginWord {
			// 反转路径，因为我们是从终点向起点回溯
			reversedPath := make([]string, len(path))
			copy(reversedPath, path)
			for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
				reversedPath[i], reversedPath[j] = reversedPath[j], reversedPath[i]
			}
			reversedPath = append([]string{beginWord}, reversedPath...)
			result = append(result, reversedPath)
			return
		}

		for _, parent := range parents[currentWord] {
			backtrack(parent, append(path, currentWord))
		}
	}

	backtrack(endWord, []string{})

	return result
}
