/*
 * @lc app=leetcode.cn id=173 lang=golang
 *
 * [173] 二叉搜索树迭代器
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
type BSTIterator struct {
	cur *TreeNode
	s   []*TreeNode
}

// 最简单的方法：用中序遍历把节点存起来
// 方法2：使用栈
func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{cur: root}
}

func (this *BSTIterator) Next() int {
	for this.cur != nil {
		this.s = append(this.s, this.cur)
		this.cur = this.cur.Left
	}
	this.cur = this.s[len(this.s)-1]
	this.s = this.s[:len(this.s)-1]
	val := this.cur.Val
	this.cur = this.cur.Right
	return val
}

func (this *BSTIterator) HasNext() bool {
	return this.cur != nil || len(this.s) != 0

}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end

