/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	// 更好的方法为：先变成环，在环中循环，再拆环
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	iter := head
	count := 1
	for iter.Next != nil {
		iter = iter.Next
		count++
	}
	// iter指向尾节点，形成环
	iter.Next = head
	add := count - k%count
	for add > 0 {
		iter = iter.Next
		add--
	}
	next := iter.Next
	iter.Next = nil
	return next
}

func rotateCircle(head *ListNode, k int) *ListNode {

	// 模拟法
	if head == nil {
		return nil
	}
	if k == 0 {
		return head
	}
	listLength := length(head)
	fmt.Printf("len %d", listLength)
	k = k % listLength
	for i := 0; i < k; i++ {
		head = rotateOnce(head)
	}
	return head
}

func length(head *ListNode) int {
	if head == nil {
		return 0
	}
	count := 0
	ptr := head
	for ptr != nil {
		ptr = ptr.Next
		count++
	}
	return count
}

func rotateOnce(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	node := ListNode{
		Next: head,
	}
	tail := head
	pre := &node
	for tail.Next != nil {
		pre = tail
		tail = tail.Next
	}
	pre.Next = nil
	tail.Next = head
	node.Next = tail
	return node.Next
}

// @lc code=end

