package main

import (
	"container/list"
)

type LRUCache struct {
	MaxEntries int
	ll         *list.List
	cache      map[interface{}]*list.Element
}

type entry struct {
	key int
	val int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		MaxEntries: capacity,
		ll:         list.New(),
		cache:      make(map[interface{}]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	if this.cache == nil {
		return -1
	}
	if val, ok := this.cache[key]; ok {
		this.ll.MoveToFront(val)
		return val.Value.(*entry).val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.cache == nil {
		this.cache = make(map[interface{}]*list.Element)
		this.ll = list.New()
	}
	if val, ok := this.cache[key]; ok {
		this.ll.MoveToFront(val)
		val.Value.(*entry).val = value
		return
	}
	ele := this.ll.PushFront(&entry{key, value})
	this.cache[key] = ele
	if this.MaxEntries != 0 && this.ll.Len() > this.MaxEntries {
		// remove the last
		el := this.ll.Back()
		this.ll.Remove(el)
		delete(this.cache, el.Value.(*entry).key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
