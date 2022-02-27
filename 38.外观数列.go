/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 外观数列
 */
package main

import "fmt"

// @lc code=start
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	ans := ""
	last := countAndSay(n - 1)
	cur := last[0] - '0'
	count := 1
	for i := 1; i < len(last); i++ {
		if last[i]-'0' == cur {
			count++
		} else {
			ans += fmt.Sprintf("%d%d", count, cur)
			cur = last[i] - '0'
			count = 1
		}
	}
	ans += fmt.Sprintf("%d%d", count, cur)
	return ans
}

// @lc code=end
