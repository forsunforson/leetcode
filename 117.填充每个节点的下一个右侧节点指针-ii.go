/*
 * @lc app=leetcode.cn id=117 lang=golang
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	start := root
	for start != nil {
		var newStart, end *Node
		handle := func(cur *Node) {
			if cur == nil {
				return
			}
			if newStart == nil {
				newStart = cur
			}
			if end != nil {
				end.Next = cur
			}
			end = cur
		}
		for p := start; p != nil; p = p.Next {
			handle(p.Left)
			handle(p.Right)
		}
		start = newStart
	}
	return root
}

// @lc code=end

