/*
 * @lc app=leetcode.cn id=168 lang=golang
 *
 * [168] Excel表列名称
 */
package main

// @lc code=start
func convertToTitle(columnNumber int) string {
	ans := []byte{}
	for columnNumber > 0 {
		columnNumber--
		m := columnNumber % 26
		ans = append(ans, byte(m)+'A')
		columnNumber /= 26
	}
	ret := ""
	for i := len(ans) - 1; i >= 0; i-- {
		ret += string(ans[i])
	}
	return ret
}

// @lc code=end
