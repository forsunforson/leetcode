/*
 * @lc app=leetcode.cn id=692 lang=golang
 *
 * [692] 前K个高频单词
 */
package main

import "container/heap"

// @lc code=start
type word struct {
	key string
	val int
}
type myHeap []word

func (a myHeap) Len() int      { return len(a) }
func (a myHeap) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a myHeap) Less(i, j int) bool {
	if a[i].val != a[j].val {
		return a[i].val < a[j].val
	}
	return a[i].key > a[j].key
}
func (a *myHeap) Pop() interface{} {
	w := (*a)[len(*a)-1]
	*a = (*a)[:len(*a)-1]
	return w
}
func (a *myHeap) Push(x interface{}) {
	*a = append(*a, x.(word))
}
func topKFrequent(words []string, k int) []string {
	wordMap := map[string]int{}
	for _, w := range words {
		wordMap[w]++
	}
	hp := myHeap{}
	for key, value := range wordMap {
		heap.Push(&hp, word{key, value})
		if len(hp) > k {
			heap.Pop(&hp)
		}
	}
	ans := make([]string, k)
	idx := k - 1
	for len(hp) > 0 && idx >= 0 {
		ans[idx] = heap.Pop(&hp).(word).key
		idx--
	}
	return ans
}

// @lc code=end
