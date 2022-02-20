/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */
package main

import (
	"container/heap"
)

// @lc code=start
type counter struct {
	num   int
	value int
}

type myHeap []counter

func (a myHeap) Len() int           { return len(a) }
func (a myHeap) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a myHeap) Less(i, j int) bool { return a[i].value < a[j].value }
func (a *myHeap) Pop() interface{} {
	top := (*a)[len(*a)-1]
	*a = (*a)[:len(*a)-1]
	return top
}
func (a *myHeap) Push(x interface{}) {
	(*a) = append((*a), x.(counter))
}

func topKFrequent(nums []int, k int) []int {
	count := map[int]int{}
	for _, num := range nums {
		count[num]++
	}
	ans := []int{}
	hp := myHeap{}
	for num, c := range count {
		heap.Push(&hp, counter{num, c})

		if len(hp) > k {
			heap.Pop(&hp)
		}
	}
	for len(hp) > 0 {
		c := heap.Pop(&hp).(counter)
		ans = append(ans, c.num)
	}
	return ans
}

// @lc code=end
