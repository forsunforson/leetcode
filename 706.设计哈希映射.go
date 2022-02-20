/*
 * @lc app=leetcode.cn id=706 lang=golang
 *
 * [706] 设计哈希映射
 */
package main

// @lc code=start
type MyHashMap struct {
	arr []int
}

func Constructor() MyHashMap {
	m := MyHashMap{arr: make([]int, 1e6+1)}
	for i := range m.arr {
		m.arr[i] = -1
	}
	return m
}

func (this *MyHashMap) Put(key int, value int) {
	this.arr[key] = value
}

func (this *MyHashMap) Get(key int) int {
	return this.arr[key]
}

func (this *MyHashMap) Remove(key int) {
	this.arr[key] = -1
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
// @lc code=end
