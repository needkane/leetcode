/*
运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 Get 和 写入数据 Put 。

获取数据 Get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
写入数据 Put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。

进阶:

你是否可以在 O(1) 时间复杂度内完成这两种操作？

示例:

LRUCache cache = new LRUCache(  缓存容量2  );

cache.Put(1, 1);
cache.Put(2, 2);
cache.Get(1);       // 返回  1
cache.Put(3, 3);    // 该操作会使得密钥 2 作废
cache.Get(2);       // 返回 -1 (未找到)
cache.Put(4, 4);    // 该操作会使得密钥 1 作废
cache.Get(1);       // 返回 -1 (未找到)
cache.Get(3);       // 返回  3
cache.Get(4);       // 返回  4


*/
package main

import (
	"container/list"
	"testing"

	"github.com/stretchr/testify/assert"
)

type LRUCache struct {
	Capacity  int
	cacheList *list.List
	cacheMap  map[interface{}]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity:  capacity,
		cacheList: list.New(),
		cacheMap:  make(map[interface{}]*list.Element),
	}
}

type entry struct {
	key   interface{}
	value interface{}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.cacheMap[key]
	if ok {
		this.cacheList.MoveToFront(v)
		return v.Value.(*entry).value.(int)
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {

	v, ok := this.cacheMap[key]
	if ok {
		this.cacheList.MoveToFront(v)
		v.Value.(*entry).value = value
	} else {
		ele := this.cacheList.PushFront(&entry{key, value})
		this.cacheMap[key] = ele
		if this.cacheList.Len() > this.Capacity {
			lastEle := this.cacheList.Back()
			delete(this.cacheMap, lastEle.Value.(*entry).key)
			this.cacheList.Remove(lastEle)
		}

	}

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	var t *testing.T
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	assert.Equal(t, cache.Get(1), 1)  // 返回  1
	cache.Put(3, 3)                   // 该操作会使得密钥 2 作废
	assert.Equal(t, cache.Get(2), -1) // 返回 -1 (未找到)
	cache.Put(4, 4)                   // 该操作会使得密钥 1 作废
	assert.Equal(t, cache.Get(1), -1) // 返回 -1 (未找到)
	assert.Equal(t, cache.Get(3), 3)  // 返回  3
	assert.Equal(t, cache.Get(4), 4)  // 返回  4
}
