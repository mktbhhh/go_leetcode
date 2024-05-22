package main

import "fmt"

type LRUCache struct {
	capacity   int
	size       int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func (l *LRUCache) addToHead(node *DLinkedNode) {
	if l.size == 0 {
		l.head = node
		l.tail = node
		l.head.next = l.tail
		l.tail.prev = l.head
	} else {
		node.prev = l.head.prev
		l.head.prev = node
		node.next = l.head
		l.head = node
		l.tail.next = l.head
	}
}

func (l *LRUCache) removeNode(node *DLinkedNode) {
	node.next.prev = node.prev
	node.prev.next = node.next
	if node == l.tail {
		l.tail = l.tail.prev
	}
	if node == l.head {
		l.head = l.head.next
	}
}

func (l *LRUCache) moveToHead(node *DLinkedNode) {
	if l.size == 0 {
		l.addToHead(node)
	} else {
		l.removeNode(node)
		l.addToHead(node)
	}
}

func (l *LRUCache) removeTail() *DLinkedNode {
	node := l.tail
	l.removeNode(l.tail)
	return node
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		capacity: capacity,
		cache:    map[int]*DLinkedNode{},
	}

	return cache
}

func (l *LRUCache) Get(key int) int {
	if _, ok := l.cache[key]; !ok {
		return -1
	}

	node := l.cache[key]
	l.moveToHead(node)
	return node.value
}

func (l *LRUCache) Put(key int, value int) {
	if _, ok := l.cache[key]; !ok {
		node := &DLinkedNode{
			key:   key,
			value: value,
		}
		l.cache[key] = node
		l.addToHead(node)
		l.size++
		if l.size > l.capacity {
			tail := l.removeTail()
			delete(l.cache, tail.key)
			l.size--
		}
	} else {
		node := l.cache[key]
		node.value = value
		l.moveToHead(node)
	}
}

func main() {
	l := Constructor(2)
	l.Put(1, 1)
	l.Put(2, 2)
	fmt.Println(l.Get(1))
	l.Put(3, 3)
	fmt.Println(l.Get(2))
	l.Put(4, 4)
	fmt.Println(l.Get(1))
	fmt.Println(l.Get(3))
	fmt.Println(l.Get(4))
}
