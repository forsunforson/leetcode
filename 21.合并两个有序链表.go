/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */
package main

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
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	// l1头节点小于l2头节点
	if l2.Val < l1.Val {
		return mergeTwoLists(l2, l1)
	}

	tail := l1
	l1Iter, l2Iter := l1.Next, l2
	for l1Iter != nil && l2Iter != nil {
		if l1Iter.Val < l2Iter.Val {
			tail.Next = l1Iter
			tail = tail.Next
			l1Iter = l1Iter.Next
		} else {
			tail.Next = l2Iter
			tail = tail.Next
			l2Iter = l2Iter.Next
		}
	}
	if l1Iter == nil {
		tail.Next = l2Iter
	} else {
		tail.Next = l1Iter
	}

	return l1
}

// @lc code=end
