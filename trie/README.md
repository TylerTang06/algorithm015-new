# 学习笔记

## Trie树

- 字典树，又称单词查找树或者健树，是一种树形结构(多叉树)
- 典型应用于统计和排序大量的字符串
- 经常被搜索引擎系统用于文本词频统计
- 优点：最大限度地减少了无谓字符串比较，查找效率比哈希表高
- 空间换时间的思想，利用字符串的公共前缀来降低查询的开销

### Trie树基本性质

- 结点本身不存完整单词
- 从根结点到某一结点，路径上经过的字符连起来，为结点对应的字符串
- 每个结点的所有子结点路径代表的字符都不相同

### Trie树相比hash表

- 可以方便找到具有同一前缀的全部键值
- 按词典序枚举字符串的数据集

### Trie树实现

```golang
type Trie struct {
	next  [26]*Trie
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	trie := this
	for _, v := range word {
		if trie.next[v-'a'] == nil {
			trie.next[v-'a'] = &Trie{}
		}
		trie = trie.next[v-'a']
	}
	trie.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	trie := this
	for _, v := range word {
		if trie.next[v-'a'] == nil {
			return false
		}
		trie = trie.next[v-'a']
	}

	return trie.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	trie := this
	for _, v := range prefix {
		if trie.next[v-'a'] == nil {
			return false
		}
		trie = trie.next[v-'a']
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
```

## 并查集

- 使用场景：找朋友类似的组团、配对的问题
- 比如，判断两个个体是不是属于同个集合

### 并查集的实现

```golang
type UnionSet struct {
    count int
    parent []int
}

func NewUnionSet(n int) *UnionSet {
    u := &UnionSet{count: n, parent: make([]int, n)}
    for i := 0; i < n; i++ {
        u.parent[i] = i
    }

    return u
}

func (u *UnionSet)Find(p int) int {
    for p != u.parent[p] {
        u.parent[p] = u.parent[u.parent[p]]
        p = u.parent[p]
    }

    return p
}

func (u *UnionSet)Union(p, q int) {
    pP := u.Find(p)
    qP := u.Find(q)
    if pP == qP {
        return 
    }

    u.parent[pP] = qP
    u.count--
}

func (u *UnionSet)Count() int {
    return u.count
}
```

## 高级搜索

- 剪枝

```golang
// 八皇后剪枝
if _, ok := colMap[col]; ok {
    continue
}
if _, ok := sumMap[row+col]; ok {
    continue
}
if _, ok := diffMap[row-col]; ok {
    continue
}

// 生成括号剪枝
if left < n {
    res = generateParenthesisRec(n, left+1, right, cur+"(", res)
}
if right < left {
    res = generateParenthesisRec(n, left, right+1, cur+")", res)
}

// 数独剪枝
func isValidSudokuOne(board [][]byte, row, col int, val byte) bool {
	for i := 0; i < 9; i++ {
		if i != row && board[i][col] != '.' && board[i][col] == val {
			return false
		}
		if i != col && board[row][i] != '.' && board[row][i] == val {
			return false
		}

		x := 3*(row/3) + i/3
		y := 3*(col/3) + i%3
		if x == row && y == col {
			continue
		}
		if board[x][y] != '.' && board[x][y] == val {
			return false
		}
	}

	return true
}
```

- 双向BFS

```golang
// 单词接龙 使用双向BFS
// 对比BFS，BFS就像 1传2，2传4一样，情况越来越多
// 而双向BFS，就像两个 1传2就找到解了
// 因此双向BFS节省了很多时间
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
```

- 启发式搜索

## 红黑树和AVL树

### 红黑树

- 每个结点不是红色就是黑色
- 根结点为黑色
- 如果结点为红，其子结点必须为黑
- 任意结点至NULL的任何路径，所包含的黑色结点数必须相同 

### AVL树

- 平衡因子：它的左子树的高度减去右子树的高度 = {-1, 0, 1}
- 不足：结点需要存储额外的信息，且调整次数频繁
- 四种旋转：左旋，右旋，左右旋，右左旋

## 总结

- Trie树和并查集用来解决特定场景的问题
- 对高级搜索中双向BFS印象最为深刻，因为对查找单词使用双向BFS后，节省了很多时间
- 解题时，第一时刻总是不能想到很好的解法，而且容易忘记做过的题，需反复练习