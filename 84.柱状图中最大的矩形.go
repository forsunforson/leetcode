/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */
package main

// @lc code=start
func largestRectangleArea(heights []int) int {
	// 利用单调栈的特性，选择出每个矩形左右最低高度
	maxSquare := 0
	left := make([]int, len(heights)) // 用来记录左右边界
	right := make([]int, len(heights))
	var leftStack, rightStack []int // 左右边界单调栈
	for i := 0; i < len(heights); i++ {
		left[i] = pushStack(&leftStack, heights, i, false)
		right[len(heights)-i-1] = pushStack(&rightStack, heights, len(heights)-1-i, true)
		//fmt.Printf("%v\n", leftStack)
	}
	for i := 0; i < len(heights); i++ {
		//fmt.Printf("idx: %d, left: %d, right: %d\n", i, left[i], right[i])
		square := (right[i] - left[i] - 1) * heights[i]
		if square > maxSquare {
			maxSquare = square
		}
	}
	return maxSquare
}

// 当前矩形入栈，返回边界
func pushStack(stack *[]int, heights []int, idx int, reversed bool) int {
	for len(*stack) > 0 {
		if heights[(*stack)[len(*stack)-1]] >= heights[idx] {
			*stack = (*stack)[:len(*stack)-1]
		} else {
			break
		}
	}
	*stack = append(*stack, idx)
	if len(*stack) == 1 {
		if reversed {
			return len(heights)
		} else {
			return -1
		}
	} else {
		return (*stack)[len(*stack)-2]
	}
}
func largestRectangleAreaOn(heights []int) int {
	// 以当前柱为起点，往后遍历，高度由最矮的那根决定
	maxSquare := 0
	for idx := range heights {
		cur := curMax(heights, idx)
		if maxSquare < cur {
			maxSquare = cur
		}
	}
	return maxSquare
}

func curMax(heights []int, curIdx int) int {
	minHeight := heights[curIdx]
	maxSquare := 0
	for i := curIdx; i < len(heights); i++ {
		if heights[i] < minHeight {
			minHeight = heights[i]
		}
		cur := (i - curIdx + 1) * minHeight
		if cur > maxSquare {
			maxSquare = cur
		}
	}
	return maxSquare
}

// @lc code=end
