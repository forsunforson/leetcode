/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}
	if left == right {
		return head
	}
	node := &ListNode{
		Next: head,
	}
	pre := node
	next := node
	lNode, rNode := node, node
	step := 0
	for step < right {
		if step < left {
			pre = lNode
			lNode = lNode.Next
		}
		rNode = rNode.Next
		next = rNode.Next
		step++
	}
	h, t := reverse(lNode, rNode)
	pre.Next = h
	t.Next = next
	return node.Next
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	cur := head
	pre := head
	for pre != tail {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return tail, head
}

// @lc code=end

