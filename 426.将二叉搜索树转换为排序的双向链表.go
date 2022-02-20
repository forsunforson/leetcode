package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func convBST2List(root *TreeNode) (*TreeNode, *TreeNode) {
	if root == nil {
		return nil, nil
	}
	var lHead, lTail *TreeNode
	if root.Left == nil {
		lHead = root
	} else {
		lHead, lTail = convBST2List(root.Left)
		lTail.Right = root
	}
	var rHead, rTail *TreeNode
	if root.Right == nil {
		rTail = root
	} else {
		rHead, rTail = convBST2List(root.Right)
		rHead.Left = root
	}
	root.Left = lTail
	root.Right = rHead
	return lHead, rTail
}

func main() {
	nodes := make([]*TreeNode, 0)
	for i := 0; i < 5; i++ {
		nodes = append(nodes, &TreeNode{
			Val: i,
		})
	}
	nodes[2].Left = nodes[0]
	nodes[2].Right = nodes[4]
	nodes[0].Right = nodes[1]
	nodes[4].Left = nodes[3]
	head, tail := convBST2List(nodes[2])
	for head != tail {
		fmt.Printf("%d\n", head.Val)
		head = head.Right
	}
}
