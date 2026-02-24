package week07

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
