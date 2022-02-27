/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */
package main

import "strings"

// @lc code=start
func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	strs := strings.Split(s, " ")
	return len(strs[len(strs)-1])
}

// @lc code=end
