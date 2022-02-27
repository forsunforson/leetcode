/*
 * @lc app=leetcode.cn id=135 lang=golang
 *
 * [135] 分发糖果
 */
package main

// @lc code=start
func candy(ratings []int) int {
	// 贪心？
	// 最低分的孩子1颗，如果向左向右递推
	left, right := make([]int, len(ratings)), make([]int, len(ratings))
	for i := range ratings {
		if i > 0 {
			if ratings[i] > ratings[i-1] {
				left[i] = left[i-1] + 1
			} else {
				left[i] = 1
			}
		} else {
			left[i] = 1
		}
	}
	for i := len(ratings) - 1; i >= 0; i-- {
		if i < len(ratings)-1 {
			if ratings[i] > ratings[i+1] {
				right[i] = right[i+1] + 1
			} else {
				right[i] = 1
			}
		} else {
			right[i] = 1
		}
	}
	sum := 0
	for i := range left {
		sum += max(left[i], right[i])
	}
	return sum
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// @lc code=end
