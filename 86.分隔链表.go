/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 */
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	littleHead := &ListNode{}
	littleCur := littleHead
	bigHead := &ListNode{}
	bigCur := bigHead
	for cur := head; cur != nil; cur = cur.Next {
		if cur.Val < x {
			littleCur.Next = cur
			littleCur = littleCur.Next
		} else {
			bigCur.Next = cur
			bigCur = bigCur.Next
		}
	}
	littleCur.Next = nil
	// for cur := littleHead; cur != nil; cur = cur.Next {
	// 	fmt.Printf("%d ", cur.Val)
	// }
	// fmt.Println()
	bigCur.Next = nil
	// for cur := bigHead; cur != nil; cur = cur.Next {
	// 	fmt.Printf("%d ", cur.Val)
	// }
	// fmt.Println()
	littleCur.Next = bigHead.Next
	return littleHead.Next
}

// @lc code=end
