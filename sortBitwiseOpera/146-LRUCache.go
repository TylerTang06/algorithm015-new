package week08

import (
	"container/list"
)

type LRUCache struct {
	cache *list.List
	val   map[int]*list.Element
	cap   int
}

func Constructor(capacity int) LRUCache {
	l := list.New()
	return LRUCache{cache: l, cap: capacity, val: make(map[int]*list.Element, capacity)}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.val[key]
	if !ok {
		return -1
	}

	this.cache.Remove(v)
	val := v.Value.([2]int)
	this.cache.PushBack(val)
	this.val[key] = this.cache.Back()
	return val[1]
}

func (this *LRUCache) Put(key int, value int) {
	if val := this.Get(key); val == -1 || val != value {
		if val != -1 && val != value {
			this.cache.Remove(this.val[key])
		}
		if this.cache.Len() == this.cap {
			v := this.cache.Front()
			this.cache.Remove(v)
			delete(this.val, v.Value.([2]int)[0])
		}
		this.cache.PushBack([2]int{key, value})
		this.val[key] = this.cache.Back()
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
