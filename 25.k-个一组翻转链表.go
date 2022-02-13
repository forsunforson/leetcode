/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	node := &ListNode{Next: head}
	pre := node
	// 反转局部链表
	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			// 如果长度不足k，则直接返回头节点
			tail = tail.Next
			if tail == nil {
				return node.Next
			}
		}
		// 防止丢失尾节点的下一个节点
		next := tail.Next
		// 局部反转
		head, tail = reverseLink(head, tail)
		// 将反转完的链表接回
		pre.Next = head
		tail.Next = next
		pre = tail
		head = next
	}
	return node.Next
}

func reverseLink(head, tail *ListNode) (*ListNode, *ListNode) {
	cur := head
	pre := tail.Next
	// 判断条件是pre与尾节点重合
	// 因为pre存的是当前反转节点的前一个，当pre == tail时，表示尾节点已经反转完了
	for pre != tail {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return tail, head
}

// @lc code=end

