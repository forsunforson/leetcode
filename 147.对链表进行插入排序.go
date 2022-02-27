/*
 * @lc app=leetcode.cn id=147 lang=golang
 *
 * [147] 对链表进行插入排序
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{
		Next: head,
	}
	cur := head.Next
	head.Next = nil
	for cur != nil {
		ptr := dummy.Next
		pre := dummy
		for ptr != nil && ptr.Val < cur.Val {
			pre = ptr
			ptr = ptr.Next
		}
		pre.Next = cur
		next := cur.Next
		cur.Next = ptr
		cur = next
	}

	return dummy.Next
}

// @lc code=end

