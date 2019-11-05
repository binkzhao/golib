package hash

import (
	"fmt"
	"sync"
)

type Key interface{}

type Value interface{}

// 哈希表结构
type HashTable struct {
	items map[int]Value
	lock  sync.RWMutex
}

func (ht *HashTable) Put(key Key, value Value) {
	ht.lock.Lock()
	defer ht.lock.Unlock()
	if ht.items == nil {
		ht.items = make(map[int]Value)
	}
	hash := ht.hash(key)
	ht.items[hash] = value
}

func (ht *HashTable) Get(key Key, value Value) Value {
	ht.lock.RLock()
	defer ht.lock.RUnlock()
	hash := ht.hash(key)
	return ht.items[hash]
}

func (ht *HashTable) Remove(key Key) {
	ht.lock.Lock()
	defer ht.lock.Unlock()
	hash := ht.hash(key)
	delete(ht.items, hash)
}

func (ht *HashTable) Size() int {
	ht.lock.RLock()
	defer ht.lock.RUnlock()
	return len(ht.items)
}

// 使用霍纳规则在 O(n) 复杂度内生成 key 的哈希值
func (HashTable) hash(key Key) int {
	k := fmt.Sprintf("%s", key)
	hash := 0
	for i := 0; i < len(k); i++ {
		hash = 31*hash + int(k[i])
	}
	return hash
}
