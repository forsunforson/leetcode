/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 先遍历两个链表获得长度
	if headA == nil || headB == nil {
		return nil
	}
	m, n := 1, 1
	n1, n2 := headA, headB
	for n1 != nil {
		m++
		n1 = n1.Next
	}
	for n2 != nil {
		n++
		n2 = n2.Next
	}
	if m < n {
		m, n = n, m
		n1, n2 = headB, headA
	} else {
		n1, n2 = headA, headB
	}
	for i := 0; i < m-n; i++ {
		n1 = n1.Next
	}
	for n1 != nil && n2 != nil {
		if n1 == n2 {
			return n1
		}
		n1 = n1.Next
		n2 = n2.Next
	}
	return nil
}

// @lc code=end

