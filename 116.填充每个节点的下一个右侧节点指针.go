/*
 * @lc app=leetcode.cn id=116 lang=golang
 *
 * [116] 填充每个节点的下一个右侧节点指针
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
// 类似于层次遍历
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		root.Left.Next = root.Right
	}
	if root.Right != nil && root.Next != nil {
		root.Right.Next = root.Next.Left
	}
	connect(root.Left)
	connect(root.Right)
	return root
}

// @lc code=end
func level(root *Node) *Node {
	if root == nil {
		return nil
	}
	end := &Node{}
	q := []*Node{root, end}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur == end {
			continue
		}
		if cur.Left != nil {
			q = append(q, cur.Left)
		}
		if cur.Right != nil {
			q = append(q, cur.Right)
		}
		if q[0] == end {
			cur.Next = nil
			q = append(q, end)
		} else {
			cur.Next = q[0]
		}
	}
	return root
}
