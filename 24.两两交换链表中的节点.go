/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	node := ListNode{
		Next: head,
	}
	pre := &node
	for head != nil {
		tail := pre
		if tail.Next != nil && tail.Next.Next != nil {
			tail = tail.Next.Next
		} else {
			return node.Next
		}

		next := tail.Next
		pre.Next = tail
		tail.Next = head
		head.Next = next
		pre = head
		head = next
	}
	return node.Next
}

// @lc code=end

