/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	addFlag := 0
	head := new(ListNode)
	var cur *ListNode = head
	for l1 != nil || l2 != nil {
		var l1Val, l2Val int
		if l1 == nil {
			l1Val = 0
		} else {
			l1Val = l1.Val
		}
		if l2 == nil {
			l2Val = 0
		} else {
			l2Val = l2.Val
		}
		num := l1Val + l2Val + addFlag
		addFlag = 0
		if num >= 10 {
			num = num - 10
			addFlag = 1
		}
		newNode := &ListNode{
			Val:  num,
			Next: nil,
		}
		cur.Next = newNode
		cur = cur.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if addFlag == 1 {
		cur.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	}
	return head.Next
}

// @lc code=end
