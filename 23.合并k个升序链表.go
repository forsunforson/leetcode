/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */
package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type NodeHeap []*ListNode

func (n NodeHeap) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
func (n NodeHeap) Less(i, j int) bool {
	return n[i].Val < n[j].Val
}
func (n NodeHeap) Len() int {
	return len(n)
}
func (n *NodeHeap) Pop() interface{} {
	top := (*n)[len(*n)-1]
	*n = (*n)[0 : len(*n)-1]
	return top
}
func (n *NodeHeap) Push(x interface{}) {
	*n = append(*n, x.(*ListNode))
}

// 用小顶堆来获得当前最小的值
func mergeKLists(lists []*ListNode) *ListNode {
	minHeap := NodeHeap{}
	heap.Init(&minHeap)
	for _, node := range lists {
		if node == nil {
			continue
		}
		heap.Push(&minHeap, node)
	}
	fmt.Printf("heap: %v \n", minHeap)
	head := ListNode{}
	cur := &head
	for {
		if minHeap.Len() < 1 {
			break
		}

		top := heap.Pop(&minHeap)
		node := top.(*ListNode)
		cur.Next = node
		cur = cur.Next
		if node.Next != nil {
			heap.Push(&minHeap, node.Next)
		}
	}

	return head.Next
}

// @lc code=end
