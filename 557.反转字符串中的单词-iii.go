/*
 * @lc app=leetcode.cn id=557 lang=golang
 *
 * [557] 反转字符串中的单词 III
 */
package main

// @lc code=start
func reverseWords(s string) string {
	strByte := []byte(s)
	l, r := 0, 0
	for r < len(strByte) {
		for r < len(strByte) && strByte[r] != ' ' {
			r++
		}
		reverseHelper(&strByte, l, r)
		l = r + 1
		r = r + 1
	}
	return string(strByte)
}

func reverseHelper(str *[]byte, l, r int) {
	for l < r {
		(*str)[l], (*str)[r-1] = (*str)[r-1], (*str)[l]
		l++
		r--
	}
}

// @lc code=end
