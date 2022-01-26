/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存机制
 */
package main

// @lc code=start
type listNode struct {
	Key  int
	Val  int
	Pre  *listNode
	Next *listNode
}

type LRUCache struct {
	cache      map[int]*listNode
	head, tail *listNode
	used       int
	capacity   int
}

func Constructor(capacity int) LRUCache {
	head := &listNode{}
	tail := &listNode{}
	head.Next = tail
	tail.Pre = head
	instance := LRUCache{
		cache:    make(map[int]*listNode),
		head:     head,
		tail:     tail,
		used:     0,
		capacity: capacity,
	}
	return instance
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	// 更新并返回
	pre := node.Pre
	next := node.Next
	pre.Next = next
	next.Pre = pre
	curFirst := this.head.Next
	this.head.Next = node
	node.Pre = this.head
	node.Next = curFirst
	curFirst.Pre = node
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cache[key]
	if !ok {
		// 第一次插入
		newNode := &listNode{
			Key: key,
			Val: value,
		}
		curFirst := this.head.Next
		this.head.Next = newNode
		newNode.Pre = this.head
		newNode.Next = curFirst
		curFirst.Pre = newNode
		this.cache[key] = newNode
		// 未满
		if this.used < this.capacity {
			this.used++
			return
		} else {
			// 已满，淘汰末尾
			curLast := this.tail.Pre
			delete(this.cache, curLast.Key)
			curLast.Pre.Next = this.tail
			this.tail.Pre = curLast.Pre
			return
		}
	}
	// 更新值
	node.Val = value
	pre := node.Pre
	next := node.Next
	pre.Next = next
	next.Pre = pre
	curFirst := this.head.Next
	this.head.Next = node
	node.Pre = this.head
	node.Next = curFirst
	curFirst.Pre = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
