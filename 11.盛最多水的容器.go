/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */
package main

// @lc code=start
func maxArea(height []int) int {
	n := len(height)
	maxSquare := 0
	l, r := 0, n-1
	for {
		if l == r {
			break
		}
		h := height[l]
		moveLeft := true
		if height[r] < h {
			h = height[r]
			moveLeft = false
		}
		square := h * (r - l)
		if maxSquare < square {
			maxSquare = square
		}
		if moveLeft {
			l++
		} else {
			r--
		}
	}
	return maxSquare
}

// @lc code=end
