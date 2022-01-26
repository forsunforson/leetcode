/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */
package main

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for {
		if fast != nil && fast.Next != nil {
			fast = fast.Next.Next
		} else {
			return nil
		}
		slow = slow.Next
		if slow == fast {
			break
		}
	}
	fast = head
	for {
		if fast == slow {
			return fast
		}
		slow = slow.Next
		fast = fast.Next
	}
}

// @lc code=end
