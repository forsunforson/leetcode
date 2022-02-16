/*
 * @lc app=leetcode.cn id=297 lang=golang
 *
 * [297] 二叉树的序列化与反序列化
 */
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	ans := ""
	var preOrder func(*TreeNode)
	preOrder = func(root *TreeNode) {
		if root == nil {
			ans += "#,"
			return
		}
		ans += fmt.Sprintf("%d,", root.Val)
		preOrder(root.Left)
		preOrder(root.Right)
	}
	preOrder(root)
	return ans
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	nodes := strings.Split(data, ",")
	var preOrder func() *TreeNode
	preOrder = func() *TreeNode {
		if nodes[0] == "#" {
			nodes = nodes[1:]
			return nil
		}
		root := &TreeNode{}
		root.Val, _ = strconv.Atoi(nodes[0])
		nodes = nodes[1:]
		root.Left = preOrder()
		root.Right = preOrder()

		return root
	}
	root := preOrder()

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
// @lc code=end
