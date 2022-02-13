/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	node := &ListNode{
		Next: head,
	}
	slow, fast := node, node.Next
	for slow != fast {
		slow = slow.Next
		if fast == nil || fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
	}
	return true
}

// @lc code=end

