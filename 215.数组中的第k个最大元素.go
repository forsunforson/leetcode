/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */
package main

// @lc code=start
type heap struct {
	k int
	h []int
}

func (h *heap) add(i int) {
	if len(h.h) < h.k {
		h.h = append(h.h, i)
		// 排序
		h.sort()
	} else {
		if i > h.top() {
			h.h[0], h.h[h.k-1] = h.h[h.k-1], h.h[0]
			h.h[h.k-1] = i
			// 排序
			h.sort()
		}
	}
}

func (h *heap) sort() {
	for i := len(h.h) - 1; i >= len(h.h)/2; i-- {
		for j := i; j > 0; j = (j - 1) / 2 {
			if h.h[(j-1)/2] < h.h[j] {
				continue
			}
			h.h[(j-1)/2], h.h[j] = h.h[j], h.h[(j-1)/2]
		}
	}
}

func (h *heap) top() int {
	return h.h[0]
}
func findKthLargest(nums []int, k int) int {
	// sol1 partition
	// sol2 小顶堆
	h := heap{
		k: k,
		h: make([]int, 0),
	}
	for _, num := range nums {
		h.add(num)
	}
	return h.top()
}

// @lc code=end
