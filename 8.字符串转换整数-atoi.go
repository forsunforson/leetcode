/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */
package main

import "strings"

// @lc code=start
func myAtoi(s string) int {
	s = strings.TrimLeft(s, " ")
	if len(s) < 1 {
		return 0
	}
	posFlag := 1
	var num int64
	if s[0] == '+' || s[0] == '-' {
		if s[0] == '-' {
			posFlag = -1
		}
		num = parseNum(s[1:])
	} else {
		num = parseNum(s)
	}

	if posFlag == 1 && num >= 2<<30-1 {
		return 2<<30 - 1
	} else if posFlag == -1 && num >= 2<<30 {
		return -1 * 2 << 30
	}
	return posFlag * int(num)
}

func parseNum(s string) int64 {
	var curNum int64
	for _, v := range s {
		if v < '0' || v > '9' {
			return curNum
		}
		curNum = curNum*10 + int64(v-'0')
		if curNum >= 2<<30 {
			break
		}
	}
	return curNum
}

// @lc code=end
