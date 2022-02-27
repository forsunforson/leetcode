/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层序遍历 II
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
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	top := [][]int{}
	q := []*TreeNode{root}

	for len(q) > 0 {
		curAns := []int{}
		size := len(q)
		for i := 0; i < size; i++ {
			cur := q[0]
			q = q[1:]
			curAns = append(curAns, cur.Val)
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		top = append(top, curAns)
	}
	ans := [][]int{}
	for i := len(top) - 1; i >= 0; i-- {
		ans = append(ans, top[i])
	}
	return ans
}

// @lc code=end

