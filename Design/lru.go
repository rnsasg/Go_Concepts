package main

import (
	"container/list"
	"fmt"
)

func main() {
	lru := Constructor(2)
	lru.Put(1, 0)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1))
	lru.Put(3, 3)
	fmt.Println(lru.Get(2))
	lru.Put(4, 4)
	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(3))
	fmt.Println(lru.Get(4))
}

type LRUCache struct {
	capacity int
	cache    *list.List
	mp       map[int]*list.Element
	// mutex    *sync.RWMutex
}

type lruNode struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	lruList := list.New()
	// mutex := &sync.RWMutex{}
	lruMap := make(map[int]*list.Element)
	return LRUCache{
		capacity: capacity,
		cache:    lruList,
		mp:       lruMap,
		// mutex:    mutex,
	}
}

func (this *LRUCache) Get(key int) int {
	// this.mutex.RLock()
	// defer this.mutex.RUnlock()
	//fmt.Printf("GET ")
	// this.print()
	if item, ok := this.mp[key]; ok {
		this.cache.MoveToFront(item)
		// this.print()
		return item.Value.(*lruNode).value
	}

	return -1
}

func (this *LRUCache) print() {
	for e := this.cache.Front(); e != nil; e = e.Next() {
		fmt.Printf("[ %d --> %d] ", e.Value.(*lruNode).key, e.Value.(*lruNode).value)
	}
	fmt.Println("")
}

func (this *LRUCache) Put(key int, value int) {
	//fmt.Printf("PUT ")
	// this.mutex.RLock()
	// defer this.mutex.RUnlock()
	if item, ok := this.mp[key]; ok {
		this.cache.MoveToFront(item)
		item.Value.(*lruNode).value = value
		// this.print()
		return
	}

	if len(this.mp) >= this.capacity {
		backElement := this.cache.Back()
		delete(this.mp, backElement.Value.(*lruNode).value)
		this.cache.Remove(backElement)
	}

	element := this.cache.PushFront(&lruNode{key, value})
	this.mp[key] = element
	// this.print()
	//fmt.Println("PUT END")
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
