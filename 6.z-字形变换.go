/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */
package main

// @lc code=start
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	result := make([]string, numRows)
	row := 0
	flag := false
	for i := 0; i < len(s); i++ {
		result[row] += s[i : i+1]
		if row == 0 || row == numRows-1 {
			flag = !flag
		}
		if flag {
			row++
		} else {
			row--
		}
	}

	ansStr := ""
	for _, row := range result {
		ansStr += row

	}
	return ansStr
}

// @lc code=end
