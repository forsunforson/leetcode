/*
 * @lc app=leetcode.cn id=461 lang=golang
 *
 * [461] 汉明距离
 */
package main

import (
	"strconv"
)

// @lc code=start
func hammingDistance(x int, y int) int {
	if x < y {
		return hammingDistance(y, x)
	}
	bitX := strconv.FormatInt(int64(x), 2)
	bitY := strconv.FormatInt(int64(y), 2)
	length := len(bitY)
	count := 0
	for i := 1; i <= length; i++ {
		if bitX[len(bitX)-i] != bitY[len(bitY)-i] {
			count++
		}
	}
	for i := 0; i < len(bitX)-length; i++ {
		if bitX[i] != '0' {
			count++
		}
	}
	return count
}

// @lc code=end
