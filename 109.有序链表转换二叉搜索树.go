/*
 * @lc app=leetcode.cn id=109 lang=golang
 *
 * [109] 有序链表转换二叉搜索树
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{
		Next: head,
	}
	slow, fast := head, head
	pre := dummy
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		pre = slow
		slow = slow.Next
	}
	pre.Next = nil
	next := slow.Next
	slow.Next = nil
	node := &TreeNode{
		Val: slow.Val,
	}
	node.Left = sortedListToBST(dummy.Next)
	node.Right = sortedListToBST(next)
	return node
}

// @lc code=end

