/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 使用自底向上的归并
	length := 0

	for cur := head; cur != nil; cur = cur.Next {
		length++

	}

	node := &ListNode{Next: head}
	for i := 1; i < length; i <<= 1 {
		pre, cur := node, node.Next
		for cur != nil {
			// 将当前链表按照长度切割，再合并
			head1 := cur
			for j := 1; j < i && cur.Next != nil; j++ {
				cur = cur.Next
			}
			head2 := cur.Next
			cur.Next = nil
			cur = head2
			for j := 1; j < i && cur != nil && cur.Next != nil; j++ {
				cur = cur.Next
			}
			var next *ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}
			pre.Next = merge(head1, head2)
			for pre.Next != nil {
				pre = pre.Next
			}
			cur = next
		}
	}
	return node.Next
}

func merge(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return dummyHead.Next
}

// @lc code=end

// 插入排序过不了
func insert(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	node := &ListNode{}
	cur := head
	for cur != nil {
		if node.Next == nil {
			node.Next = cur
			next := cur.Next
			cur.Next = nil
			cur = next
		} else {
			ptr := node.Next
			pre := node
			for ptr != nil && ptr.Val < cur.Val {
				pre = ptr
				ptr = ptr.Next
			}
			pre.Next = cur
			next := cur.Next
			cur.Next = ptr
			cur = next
		}

	}
	return node.Next
}