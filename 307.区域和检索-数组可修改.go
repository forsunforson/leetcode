/*
 * @lc app=leetcode.cn id=307 lang=golang
 *
 * [307] 区域和检索 - 数组可修改
 */
package main

// @lc code=start
type NumArray struct {
	tree []int
	n    int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	tree := make([]int, 2*n)
	for i, j := n, 0; j < n; j++ {
		tree[i] = nums[j]
		i++
	}
	for i := n - 1; i > 0; i-- {
		tree[i] = tree[i*2] + tree[i*2+1]
	}
	return NumArray{tree: tree, n: n}
}

func (this *NumArray) Update(index int, val int) {
	idx := index + this.n
	this.tree[idx] = val

	for idx > 0 {
		l, r := idx, idx
		if l%2 == 1 {
			l--
		} else {
			r++
		}
		this.tree[idx/2] = this.tree[l] + this.tree[r]
		idx = idx / 2
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	sum := 0
	l, r := left+this.n, right+this.n
	for l <= r {
		if l%2 == 1 {
			// 如果左边界是右子，那么加上本身，移动
			sum += this.tree[l]
			l++
		}
		if r%2 == 0 {
			// 如果右边界是左子，加上本身，移动
			sum += this.tree[r]
			r--
		}
		// 对于一个连续的范围，加它们的父节点就可以
		l /= 2
		r /= 2
	}
	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
// @lc code=end
