/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	// 快慢指针，当链表是偶数时，slow为中点前一个节点
	// 当链表是奇数时，slow为中点
	// 反转后半程
	slow, fast := head, head
	for {
		if fast.Next != nil && fast.Next.Next != nil {
			fast = fast.Next.Next
		} else {
			break
		}
		slow = slow.Next
	}
	bottom := reverse(slow.Next)
	ptr := head
	for bottom != nil {
		if ptr.Val != bottom.Val {
			return false
		}
		ptr = ptr.Next
		bottom = bottom.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ptr := head
	var pre *ListNode
	for ptr != nil {
		next := ptr.Next
		ptr.Next = pre
		pre = ptr
		ptr = next
	}
	return pre
}

// @lc code=end

