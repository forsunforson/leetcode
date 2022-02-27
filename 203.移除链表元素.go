/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{}
	pre := dummy
	cur := head
	for cur != nil {
		for cur != nil && cur.Val == val {
			cur = cur.Next
		}
		pre.Next = cur
		pre = cur
		if cur != nil {
			cur = cur.Next
		}
	}
	return dummy.Next
}

// @lc code=end

