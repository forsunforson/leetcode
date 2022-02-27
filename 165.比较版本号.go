/*
 * @lc app=leetcode.cn id=165 lang=golang
 *
 * [165] 比较版本号
 */
package main

import (
	"strconv"
	"strings"
)

// @lc code=start
func compareVersion(version1 string, version2 string) int {
	v1, v2 := strings.Split(version1, "."), strings.Split(version2, ".")
	size := len(v1)
	if len(v2) > size {
		size = len(v2)
	}
	for i := 0; i < size; i++ {
		i1, i2 := 0, 0
		if i < len(v1) {
			i1, _ = strconv.Atoi(v1[i])
		}
		if i < len(v2) {
			i2, _ = strconv.Atoi(v2[i])
		}
		if i1 < i2 {
			return -1
		}
		if i1 > i2 {
			return 1
		}
	}
	return 0
}

// @lc code=end
