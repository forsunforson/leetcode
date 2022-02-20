/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	next := &TreeNode{}
	q := []*TreeNode{next}
	thisLevel := []*TreeNode{root}
	ans := make([][]int, 0)
	flag := true
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur == next {
			if len(thisLevel) == 0 {
				break
			}
			l := make([]int, 0)
			q = append(q, thisLevel...)
			for i := 0; i < len(thisLevel); i++ {
				if flag {
					l = append(l, thisLevel[i].Val)
				} else {
					l = append(l, thisLevel[len(thisLevel)-i-1].Val)
				}
			}
			thisLevel = []*TreeNode{}
			flag = !flag
			ans = append(ans, l)
			q = append(q, next)
			continue
		}
		if cur.Left != nil {
			thisLevel = append(thisLevel, cur.Left)
		}
		if cur.Right != nil {
			thisLevel = append(thisLevel, cur.Right)
		}
	}
	return ans
}

// @lc code=end

